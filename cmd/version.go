package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Google Cloud Pub/Sub Client",
	Long:  `Google Cloud Pub/Sub Client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.3.0")
	},
}
