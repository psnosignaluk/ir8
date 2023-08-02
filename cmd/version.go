package cmd

import (
	"fmt"

	"github.com/psnosignaluk/ir8/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ir8",
	Long:  `Versions. Symantics. This be ours`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ir8 version %s\n", utils.AppVersion())
	},
}
