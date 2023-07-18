package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:     "[d] daily",
	Short:   "Generate daily summary",
	Long:    `Generate daily summary`,
	Aliases: []string{"d"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)
}
