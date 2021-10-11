package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var Version = "development"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the avm version",
	Run: execVersionCmd,
}

func execVersionCmd(cmd *cobra.Command, args []string) {
	log.Infof("version: %s", Version)
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
