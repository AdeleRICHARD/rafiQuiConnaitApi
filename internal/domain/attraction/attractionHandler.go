package attraction

import (
	"embed"
	_ "embed"
	"encoding/csv"
	"log"
	"rafiQuiConnaitApi/internal/domain/model"
	"rafiQuiConnaitApi/internal/domain/producer/csvReader"
	"rafiQuiConnaitApi/internal/domain/producer/elasticsearch"
	"strings"
)

//go:embed attractions.csv
var attractionsCSV embed.FS

type AttractionHandler struct {
	esClient  *elasticsearch.EsClient
	csvReader *csvReader.CSVReader
}

func NewAttractionHandler(esClient *elasticsearch.EsClient, csvReader *csvReader.CSVReader) *AttractionHandler {
	return &AttractionHandler{
		esClient:  esClient,
		csvReader: csvReader,
	}
}

func (h *AttractionHandler) ProcessCSV() error {
	attractionsCSVRead, err := attractionsCSV.ReadFile("attractions.csv")
	if err != nil {
		return err
	}

	r := csv.NewReader(strings.NewReader(string(attractionsCSVRead)))
	lines, err := r.ReadAll()
	if err != nil {
		return err
	}

	// Skip the header line
	lines = lines[1:]

	for _, line := range lines {
		attraction := model.Attraction{
			ID:        line[0], // This assumes that the first column is the ID
			Name:      line[2],
			LandName:  line[3],
			History:   line[5],
			CreatedIn: line[6],
			Creator:   line[4],
			FunFacts:  strings.Split(line[7], ", "), // assuming fun facts are comma separated
		}

		err := h.esClient.AddAttraction(attraction)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
