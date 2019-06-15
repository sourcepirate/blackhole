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

	mapping_service := elastic.NewGetMappingService(d.client)
	mapping_service.Index(index)
	meta_maps, map_err := mapping_service.Do(context.Background())
	log.Println(meta_maps)

	if map_err != nil {
		// mapping error
		log.Fatalf("Unable to fetch mapping for : %s", index)
	}

	map_encode_error := encoder.Encode(meta_maps)

	if map_encode_error != nil {
		// mapping encode error
		log.Fatalf("Mapping encode error!!")
	}

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
			data := make(map[string]interface{})
			data["_id"] = it.Id
			data["_source"] = it.Source
			data["_doc"] = it.Type
			errored = encoder.Encode(data)
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

type Data struct {
	ID     string                 `json:"_id"`
	Doc    string                 `json:"_doc"`
	Source map[string]interface{} `json:"_source"`
}

// Exporter
type Exporter struct {
	client    *elastic.Client
	patchsize int
}

func (d *Exporter) Export(index string, decoder *json.Decoder) {

	var data Data

	start := 0
	map_data := make(map[string]interface{})
	mapping_err := decoder.Decode(&map_data)

	if mapping_err != nil {
		log.Fatalf("Decode Error for mapping")
	}

	index_service := elastic.NewIndexService(d.client)

	index_service.Index(index)

	err := decoder.Decode(&data)

	if err != nil {
		log.Println("Errored")
		log.Fatal(err)
	}
	json_data := map_data[index].(map[string]interface{})["mappings"]
	json_data = json_data.(map[string]interface{})[data.Doc]
	index_service.BodyJson(json_data)
	index_service.Type(data.Doc)

	_, creation_error := index_service.Do(context.Background())
	// Error created.
	if creation_error != nil {
		log.Fatalf("Failed creating index %s", creation_error)
	}

	for decoder.More() {

		mapper := elastic.NewBulkService(d.client)

		for i := 0; i < d.patchsize; i++ {
			if err != nil {
				if err.Error() == "unexpected EOF" {
					return
				} else {
					log.Fatal("Unable to decode data")
				}
			}
			request := elastic.NewBulkIndexRequest()
			request.Index(index)
			request.Doc(data.Source)
			request.Id(data.ID)
			mapper.Index(index).Type(data.Doc).Add(request)
			err = decoder.Decode(&data)
			start += 1
		}

		log.Printf("Updated %d documents", start)
		_, errored := mapper.Do(context.Background())
		if errored != nil {
			log.Panic(errored)
		}
	}
}

func NewExporter(url string, patchsize int) (*Exporter, error) {
	client, err := getClient(url)
	if err != nil {
		return nil, err
	}
	return &Exporter{client, patchsize}, nil
}
