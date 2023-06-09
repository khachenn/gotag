package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gotag [command]",
	Short:   "gotag - a simple CLI to semantic versioning",
	Version: "0.0.1",
	Long:    `gotag is a super fancy CLI (kidding) for semantic versioning`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 1 {
			fmt.Println("gotag version", cmd.Version)
		} else {
			fmt.Println("gotag version", cmd.Version)
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
