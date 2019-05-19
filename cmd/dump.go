package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sourcepirate/blackhole/lib"
	"github.com/spf13/cobra"
)

func dumpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "dump",
		Short: "Dump elasticsearch index to json",
		Long:  "Dump elasticsearch index to json",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			arg := args[0]
			url := args[1]
			index := args[2]

			fp, err := os.Create(arg)
			if err != nil {
				log.Fatal("Unable to create file %s", arg)
			}

			writer := bufio.NewWriter(fp)

			encoder := json.NewEncoder(writer)
			dumper, errored := lib.NewDumper(url, batchsize)

			if errored != nil {
				fmt.Println(errored)
				log.Fatal("Errored while dumping ")
			}
			defer fp.Close()
			_, derr := dumper.Dump(index, encoder)
			if derr != nil {
				log.Fatal(derr)
			}
		},
	}
}
