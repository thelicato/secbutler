package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thelicato/secbutler/pkg/runners"
)

var payloadsCmd = &cobra.Command{
	Use:   "payloads",
	Short: "Obtain and serve common payloads",
	Run: func(cmd *cobra.Command, args []string) {
		runners.ServePayloads()
	},
}

func init() {
	rootCmd.AddCommand(payloadsCmd)
}
