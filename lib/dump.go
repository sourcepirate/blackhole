package lib

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"github.com/olivere/elastic"
)

func getClient(url string) (*elastic.Client, error) {
	return elastic.NewClient(elastic.SetURL(url))
}

func getQuery(client *elastic.Client, index string, from int, size int) *elastic.SearchService {
	query := elastic.NewMatchAllQuery()
	return client.Search().Index(index).Query(query).Size(size).From(from)
}

func writeHeader(*csv.Writer, result interface{}) {

}

func writeResult(*csv.Writer, result interface{}) {

}

// Dumper ...
type Dumper struct {
	client    *elastic.Client
	patchsize int
}

// Dump ...
func (d *Dumper) Dump(index string, writer *bufio.Writer) (bool, error) {
	var start int
	start = 0
	headerwritten := false

	if err != nil {
		log.Fatal("Cannot create new file %s", outfile)
		return false, "File Error"
	}

	for {
		// loops until no data available.
		searchSerivce := getQuery(d.client, index, start, d.patchsize)
		results, errored := searchSerivce.Do(context.Background())

		
	}

	return true, nil
}
