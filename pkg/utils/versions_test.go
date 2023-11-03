package utils

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestListReleasesWithPagination(t *testing.T) {
	setup(t)

	httpmock.RegisterResponderWithQuery(
		"GET",
		repositoryUrl,
		"list-type=2&prefix=cli",
		httpmock.NewXmlResponderOrPanic(200, httpmock.File("resources/all-versions-1.xml")),
	)

	httpmock.RegisterResponderWithQuery(
		"GET",
		repositoryUrl,
		"list-type=2&prefix=cli&continuation-token=l0g4nr0y",
		httpmock.NewXmlResponderOrPanic(200, httpmock.File("resources/all-versions-2.xml")),
	)

	httpmock.RegisterResponderWithQuery(
		"GET",
		repositoryUrl,
		"list-type=2&prefix=cli&continuation-token=k3nd411r0y",
		httpmock.NewXmlResponderOrPanic(200, httpmock.File("resources/all-versions-3.xml")),
	)

	versions, err := GetAllVersions()
	require.NoError(t, err)
	expected := []string{"v3.0.0", "v2.0.0", "v1.10.0", "v1.1.11"}
	assert.Equal(t, expected, versions)
}

func TestGetLatestVersion(t *testing.T) {
	setup(t)

	httpmock.RegisterResponder(
		"GET",
		repositoryUrl,
		httpmock.NewXmlResponderOrPanic(200, httpmock.File("resources/get-latest-version.xml")),
	)

	v, err := GetLatestVersion()
	require.NoError(t, err)

	assert.Equal(t, "v2.0.0", v)
}

func TestGetBinDownloadUrlForVersion(t *testing.T) {
	v := "v1.18.0"
	goos := "darwin"
	goarch := "arm64"
	assert.Equal(t, "https://armory-cli-releases.s3.amazonaws.com/cli/v1.18.0/armory-darwin-arm64", GetBinDownloadUrlForVersion(v, goos, goarch))
}

func setup(t *testing.T) {
	httpmock.ActivateNonDefault(client.HTTPClient)
	t.Cleanup(httpmock.DeactivateAndReset)
}
