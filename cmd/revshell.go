package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thelicato/secbutler/pkg/runners"
)

var revshellCmd = &cobra.Command{
	Use:   "revshell",
	Short: "Obtain the command for a reverse shell",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GetReverseShell()
	},
}

func init() {
	rootCmd.AddCommand(revshellCmd)
}
