
package cmd

import (
	"github.com/armory/avm/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "install an Armory CLI version, if version is omitted latest will be used and it will be linked to as default",
	Run: execInstallCmd,
	Args: cobra.MaximumNArgs(1),
}

func execInstallCmd(cmd *cobra.Command, args []string) {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	useVersionAsDefault, _ := cmd.Flags().GetBool("default")
	var version string
	if len(args) == 0 {
		version = utils.GetLatestVersion()
		useVersionAsDefault = true
	} else {
		version = args[0]
	}
	log.Infof("Installing version: %s for %s-%s", version, goos, goarch)

	// Ensure ~/.avm exists
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable get home dir, err: %s", err.Error())
	}

	dir := filepath.Join(home, ".avm", "versions", version)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("unable to make dir: %s, err: %s", dir, err.Error())
	}

	path := filepath.Join(dir, "armory")
	err = downloadRelease(path, utils.GetBinDownloadUrlForVersion(version))
	if err != nil {
		log.Fatalf("Failed to download release, err: %s", err.Error())
	}
	if useVersionAsDefault {
		utils.LinkToVersion(version)
	}
}

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolP("default", "d", false, "Set version as default")
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadRelease(filepath string, url string) error {
	utils.DeleteIfExists(filepath)

	log.Infof("Downloading release: %s to %s", url, filepath)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}