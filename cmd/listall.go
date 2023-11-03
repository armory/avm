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
	Run:   execListAllCmd,
}

func execListAllCmd(cmd *cobra.Command, args []string) {
	versions, err := utils.GetAllVersions()
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, version := range versions {
		log.Infof(version)
	}
}

func init() {
	RootCmd.AddCommand(listallCmd)
}
