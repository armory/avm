package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list installed Armory CLI versions",
	Run: executeListCmd,
}

func executeListCmd(cmd *cobra.Command, args []string) {
	// Ensure ~/.avm exists
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable get home dir, err: %s", err.Error())
	}
	dir := filepath.Join(home, ".avm", "versions")
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf(err.Error())
	}

	link := filepath.Join(home, ".avm", "bin", "armory")
	current, _ := os.Readlink(link)
	if current != "" {
		parts := strings.Split(current, "/")
		if len(parts) < 3 {
			return
		}
		current = parts[len(parts) - 2]
	}
	for _, file := range files {
		if file.Name() == current {
			log.Infof("%s [default]", file.Name())
			continue
		}
		log.Infof(file.Name())
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}
