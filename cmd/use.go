package cmd

import (
	"github.com/armory/avm/pkg/utils"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "select an installed Armory CLI version to use",
	Run: executeUseCmd,
	Args: cobra.ExactArgs(1),
}

func executeUseCmd(cmd *cobra.Command, args []string) {
	version := args[0]
	utils.LinkToVersion(version)
}

func init() {
	rootCmd.AddCommand(useCmd)
}
