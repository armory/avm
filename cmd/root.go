package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"path/filepath"
	"runtime"
)

const (
	COMMAND = "armory"
)

var verboseFlag bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "avm",
	Short: "Armory Version Manager",
}

func Execute() {
	//goland:noinspection GoBoolExpressions because incorrectly detects as "always false"
	if runtime.GOOS == "windows" {
		log.Fatalf("avm only supports OS X and GNU+Linux")
	}
	// Ensure ~/.avm exists
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("unable get home dir, err: %s", err.Error())
	}

	path := filepath.Join(home, ".avm")
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalf("Unable to make ~/.avm")
	}

	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "show more details")
	RootCmd.PersistentPreRunE = configureLogging
}

func configureLogging(cmd *cobra.Command, args []string) error {
	lvl := log.InfoLevel
	if verboseFlag {
		lvl = log.DebugLevel
	}
	log.SetLevel(lvl)
	log.SetFormatter(&easy.Formatter{
		LogFormat: "%msg%\n",
	})
	return nil
}
