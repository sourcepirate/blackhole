package cmd

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/sourcepirate/blackhole/lib"
	"github.com/spf13/cobra"
)

func exportCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "export [dumpfile] [url] [indexname]",
		Short: "Export elasticsearch json to index",
		Long:  "Export elasticsearch json to index",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			arg := args[0]
			url := args[1]
			index := args[2]

			fp, err := os.Open(arg)
			if err != nil {
				log.Println("Unable to open file %s", arg)
			}

			reader := bufio.NewReader(fp)
			decoder := json.NewDecoder(reader)

			exporter, errored := lib.NewExporter(url, batchsize)

			if errored != nil {
				log.Fatal("Errored durring export")
			}

			exporter.Export(index, decoder)

		},
	}
}
