package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:     "[r] report",
	Short:   "Daily Report.",
	Long:    "Daily Report.",
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("report called")
		return nil
	},
}

func init() {
	dailyCmd.AddCommand(reportCmd)

}
