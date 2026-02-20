package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thelicato/secbutler/pkg/runners"
)

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Generate a install script for the most common cybersecurity tools",
	Run: func(cmd *cobra.Command, args []string) {
		runners.GenerateToolsInstallScript()
	},
}

func init() {
	rootCmd.AddCommand(toolsCmd)
}
