package cmd

import (
	gotag "github.com/khachenn/gotag/pkg"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release [flags]",
	Aliases: []string{"r"},
	Short:   "Release new project version",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 1 {
			if ok, _ := cmd.Flags().GetBool("major"); ok {
				gotag.UpdateVersion(gotag.SvIncMajor)
				return
			}
			if ok, _ := cmd.Flags().GetBool("minor"); ok {
				gotag.UpdateVersion(gotag.SvIncMinor)
				return
			}
			if ok, _ := cmd.Flags().GetBool("patch"); ok {
				gotag.UpdateVersion(gotag.SvIncPatch)
				return
			}
		}
		cmd.Help()
	},
}

func init() {
	releaseCmd.PersistentFlags().BoolP("major", "m", false, "major")
	releaseCmd.PersistentFlags().BoolP("minor", "n", false, "minor")
	releaseCmd.PersistentFlags().BoolP("patch", "p", false, "patch")
	rootCmd.AddCommand(releaseCmd)
}
