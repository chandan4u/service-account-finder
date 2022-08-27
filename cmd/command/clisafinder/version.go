package clisafinder

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "It return the current version",
	Long:  "safinder is cmd line tool for finding service account of organisation projects",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("V1.0.0 - Welcome | Service account finder")
	},
}
