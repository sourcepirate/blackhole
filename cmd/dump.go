package cmd

import (
	"github.com/spf13/cobra"
)

func dumpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "dump",
		Short: "Dump elasticsearch index to json",
		Long:  "Dump elasticsearch index to json",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
