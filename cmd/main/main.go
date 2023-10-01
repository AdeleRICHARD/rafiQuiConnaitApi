package main

import (
	_ "embed"
	"fmt"
	"log"

	"rafiQuiConnaitApi/internal/domain/attraction"
	"rafiQuiConnaitApi/internal/domain/producer/csvReader"
	"rafiQuiConnaitApi/internal/domain/producer/elasticsearch"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	esClient, err := elasticsearch.NewEsClient()
	if err != nil {
		log.Fatalf("Error creating ES client: %s", err)
	}

	err = esClient.CreateIndex("attractions")
	if err != nil {
		fmt.Errorf("Error creating index: %s", err)
	}

	csvReader := csvReader.NewCSVReader()

	attractionHandler := attraction.NewAttractionHandler(esClient, csvReader)
	// Process the CSV file
	err = attractionHandler.ProcessCSV()
	if err != nil {
		log.Fatalf("Error processing CSV file: %s", err)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
