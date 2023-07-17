package cmd

import (
	"fmt"

	"github.com/billylkc/get_me_rich/src"
	"github.com/spf13/cobra"
)

// dailyCmd represents the daily command
var dailyCmd = &cobra.Command{
	Use:     "[d] daily",
	Short:   "Generate daily summary",
	Long:    `Generate daily summary`,
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		a := src.Hello()
		fmt.Println(a)
		fmt.Println("daily called")
	},
}

func init() {
	rootCmd.AddCommand(dailyCmd)
}
