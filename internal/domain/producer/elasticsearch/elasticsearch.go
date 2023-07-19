package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"rafiQuiConnaitApi/internal/domain/model"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type EsClient struct {
	Client *elasticsearch.Client
}

func NewEsClient() (*EsClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return &EsClient{Client: client}, nil
}

func (es *EsClient) AddAttraction(attraction model.Attraction) error {
	jsonAttraction, err := json.Marshal(attraction)
	if err != nil {
		log.Fatalf("Error marshalling the attraction: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      "attractions",
		DocumentID: attraction.ID,
		Body:       strings.NewReader(string(jsonAttraction)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Error: %s", res.String())
	}

	return nil
}

func (es *EsClient) CreateIndex(indexName string) error {
	req := esapi.IndicesCreateRequest{
		Index: indexName,
	}

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Errorf("Error creating index: %s", res.String())
	}

	return nil
}
