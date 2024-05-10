package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jirawat-rackz/da-mock-api/data/dadata"
	"github.com/jirawat-rackz/da-mock-api/handler/dahandler"
	"github.com/jirawat-rackz/da-mock-api/pkg/httpserve"
)

func main() {
	r := httpserve.New("8080")

	r.Use(
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
		}),
	)

	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// mock data
	daData, err := dadata.NewDAData()
	if err != nil {
		log.Fatalf("Failed to create DAData: %v", err)
		panic(err)
	}

	// DAHandler
	daHandler := dahandler.NewDAHandler(daData)
	daRoute := v1.Group(dahandler.BasePath)
	{
		daRoute.GET("/coin-list", daHandler.GetCoinList)
		daRoute.GET("/coin/:id", daHandler.GetCoinByID)
	}

	if err := r.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
