package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thelicato/secbutler/pkg/runners"
)

var cheatsheetCmd = &cobra.Command{
	Use:   "cheatsheet",
	Short: "Read common cheatsheets & payloads",
	Run: func(cmd *cobra.Command, args []string) {
		runners.ShowCheatsheet()
	},
}

func init() {
	rootCmd.AddCommand(cheatsheetCmd)
}
