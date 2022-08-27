package clisafinder

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "safinder",
	Short: "safinder is cmd line tool for finding service account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome | Service Account Finder command line tool for finding service account")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
