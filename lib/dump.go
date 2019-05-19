package lib

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/olivere/elastic"
)

func getClient(url string) (*elastic.Client, error) {
	return elastic.NewSimpleClient(elastic.SetURL(url))
}

func getQuery(client *elastic.Client, index string, from int, size int) *elastic.SearchService {
	query := elastic.NewMatchAllQuery()
	return client.Search().Index(index).Query(query).Size(size).From(from)
}

// Dumper ...
type Dumper struct {
	client    *elastic.Client
	patchsize int
}

// Dump ...
func (d *Dumper) Dump(index string, encoder *json.Encoder) (bool, error) {
	var start int
	start = 0

	for {
		// loops until no data available.
		log.Printf("Querying %s from %d", index, start)
		searchSerivce := getQuery(d.client, index, start, d.patchsize)
		result, errored := searchSerivce.Do(context.Background())
		if errored != nil {
			return false, errors.New("ES Fetch Error")
		}
		start = start + len(result.Hits.Hits)
		log.Printf("Fetched record %d - %d", len(result.Hits.Hits), start)

		for _, it := range result.Hits.Hits {
			errored = encoder.Encode(it)
			if errored != nil {
				return false, errors.New("File Writter error")
			}
		}

		if int64(start) >= result.TotalHits() {
			break
		}
	}

	return true, nil
}

// NewDumper ...
func NewDumper(url string, patchsize int) (*Dumper, error) {
	client, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return &Dumper{client, patchsize}, nil
}
