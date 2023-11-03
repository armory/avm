package utils

import (
	"encoding/xml"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/hashicorp/go-retryablehttp"
	"net/url"
	"sort"
	"strings"
)

const (
	repositoryUrl = "https://armory-cli-releases.s3.amazonaws.com/"

	cliPrefix = "cli"
)

var client = retryablehttp.NewClient()

// GetAllVersions returns versions available in S3 repository, sorted as SemVers in descending order.
func GetAllVersions() ([]string, error) {
	var versions []*semver.Version
	more := true
	params := url.Values{}
	params.Set("list-type", "2")
	params.Set("prefix", cliPrefix)
	for more {
		u := buildUrl(params)
		httpRes, err := client.Get(u)
		if err != nil {
			return []string{}, err
		}

		var res listBucketResult
		if err := xml.NewDecoder(httpRes.Body).Decode(&res); err != nil {
			return []string{}, err
		}
		more = res.IsTruncated
		if res.NextContinuationToken != "" {
			params.Set("continuation-token", res.NextContinuationToken)
		}

		for _, o := range res.Contents {
			v, err := o.version()
			if err != nil {
				return []string{}, err
			}
			versions = append(versions, v)
		}
	}

	sort.Sort(sort.Reverse(semver.Collection(versions)))

	var versionStrings []string
	for _, v := range versions {
		versionStrings = append(versionStrings, v.Original())
	}
	return versionStrings, nil
}

func GetLatestVersion() (string, error) {
	all, err := GetAllVersions()
	if err != nil {
		return "", err
	}
	return all[0], err
}

func GetBinDownloadUrlForVersion(version, goos, goarch string) string {
	return fmt.Sprintf("%s%s/%s/armory-%s-%s", repositoryUrl, cliPrefix, version, goos, goarch)
}

func buildUrl(params url.Values) string {
	u, err := url.Parse(repositoryUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = params.Encode()
	return u.String()
}

type listBucketResult struct {
	IsTruncated           bool
	Contents              []object
	NextContinuationToken string
}

type object struct {
	Key string // asset key has format "cli/{version}/armory-{os}-{arch}"
}

func (o object) version() (*semver.Version, error) {
	v := strings.Split(o.Key, "/")[1]
	return semver.NewVersion(v)
}
