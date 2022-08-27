package clisafinder

import (
	"github.com/spf13/cobra"
	"service-account-finder/internal/safinder"
)

func init() {
	rootCmd.AddCommand(helmCmd)
	helmCmd.Flags().StringP("org", "o", "", "Organisation id for fetching project and SA")
	helmCmd.MarkFlagRequired("file")
}

var helmCmd = &cobra.Command{
	Use:   "saf",
	Short: "It find service group for organisation projects",
	Long:  "Go binary find service account for organisation projects",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = cmd.Flags().GetString("org")
		safinder.SAFinder()
	},
}
