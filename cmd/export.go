package cmd

import (
	"github.com/spf13/cobra"
)

func exportCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "export",
		Short: "Export elasticsearch json to index",
		Long:  "Export elasticsearch json to index",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
