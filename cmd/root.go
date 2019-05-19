package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var batchsize = 20

var rootCmd = &cobra.Command{
	Use:   "blackhole",
	Short: "Tool of choice for elasticsearch dump",
	Long:  "Tool of choice for elasticsearch dump",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Initialize() {
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(dumpCommand())
	rootCmd.AddCommand(exportCommand())

	rootCmd.PersistentFlags().IntVar(&batchsize, "batchsize", 20, "Size for each batch")
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
