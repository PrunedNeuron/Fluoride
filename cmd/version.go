package cmd

import (
	"icon-requests/config"
	"fmt"

	"github.com/spf13/cobra"
)

// Version command
func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(config.Executable + " - " + config.GitVersion)
		},
	})
}
