package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thelicato/secbutler/pkg/runners"
)

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Obtain the command to start a reverse shell listener",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GetReverseShellListener()
	},
}

func init() {
	rootCmd.AddCommand(listenerCmd)
}
