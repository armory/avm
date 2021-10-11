package cmd

import (
	"github.com/armory/avm/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listallCmd represents the listall command
var listallCmd = &cobra.Command{
	Use:   "listall",
	Short: "list available versions",
	Run: execListAllCmd,
}

func execListAllCmd(cmd *cobra.Command, args []string) {
	allReleases := utils.GetAllReleases()
	for _, release := range allReleases {
		log.Infof(*release.TagName)
	}
}

func init() {
	rootCmd.AddCommand(listallCmd)
}
