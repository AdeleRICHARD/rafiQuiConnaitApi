package main

import (
	"log"
	"rafiQuiConnaitApi/internal/domain/producer/elasticsearch"

	"github.com/gin-gonic/gin"
)

var esClient *elasticsearch.EsClient

func init() {
	var err error
	esClient, err = elasticsearch.NewEsClient()
	if err != nil {
		log.Fatalf("Failed to create ES client: %v", err)
	}

	// Créer l'index s'il n'existe pas déjà
	err = esClient.CreateIndex("attractions")
	if err != nil {
		log.Fatalf("Failed to create ES index: %v", err)
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Ajoutez plus de routes qui utilisent esClient ici...

	r.Run() // listen and serve on 0.0.0.0:8080
}
