package utils

import (
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
	log "github.com/sirupsen/logrus"
	"runtime"
)

const (
	OWNER   = "armory-io"
	REPO    = "armory-cli"
	COMMAND = "armory"
)

var client = github.NewClient(nil)
var ctx = context.Background()

func GetAllReleases() []*github.RepositoryRelease {
	opt := &github.ListOptions{
		PerPage: 10,
	}
	var allReleases []*github.RepositoryRelease
	for {
		repositoryReleases, response, err := client.Repositories.ListReleases(ctx, OWNER, REPO, opt)
		if err != nil {
			log.Fatalf(err.Error())
		}
		allReleases = append(allReleases, repositoryReleases...)
		if response.NextPage == 0 {
			break
		}
		opt.Page = response.NextPage
	}

	return allReleases
}

func GetLatestVersion() string {
	repositoryRelease, _, err := client.Repositories.GetLatestRelease(ctx, OWNER, REPO)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return *repositoryRelease.TagName
}

func GetBinDownloadUrlForVersion(version string) string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	release, _, err := client.Repositories.GetReleaseByTag(ctx, OWNER, REPO, version)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var assetToInstall *github.ReleaseAsset
	var assetNameToFind = fmt.Sprintf("%s-%s-%s", COMMAND, goos, goarch)
	for _, asset := range release.Assets {
		if *asset.Name == assetNameToFind {
			assetToInstall = asset
			break
		}
	}
	if assetToInstall == nil {
		log.Fatalf("Unable to find release asset with name: %s", assetNameToFind)
	}
	return *assetToInstall.BrowserDownloadURL
}
