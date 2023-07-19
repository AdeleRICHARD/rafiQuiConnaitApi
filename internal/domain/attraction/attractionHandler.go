package attraction

import (
	"log"
	"rafiQuiConnaitApi/internal/domain/model"
	"rafiQuiConnaitApi/internal/domain/producer/csvReader"
	"rafiQuiConnaitApi/internal/domain/producer/elasticsearch"
	"strings"
)

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

func (h *AttractionHandler) ProcessCSV(file string) error {
	lines, err := h.csvReader.Read(file)
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
