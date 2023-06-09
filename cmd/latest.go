package cmd

import (
	"fmt"

	gotag "github.com/khachenn/gotag/pkg"
	"github.com/spf13/cobra"
)

var latestCmd = &cobra.Command{
	Use:                   "latest",
	DisableFlagsInUseLine: true,
	DisableFlagParsing:    true,
	Short:                 "Show latest version",
	Run: func(cmd *cobra.Command, args []string) {
		latestVersion := gotag.GetLatestVersion()
		fmt.Printf("latest: %s\n", latestVersion)
	},
}

func init() {
	rootCmd.AddCommand(latestCmd)
}
