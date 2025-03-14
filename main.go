package main

import (
	"goStreaming/pkg/api/stream/interfaces/controllers"
	"goStreaming/pkg/api/stream/interfaces/repositories"
	"goStreaming/pkg/api/stream/interfaces/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo, err := repositories.NewStreamRepository()
	if err != nil {
		log.Fatal(err)
	}

	service := services.NewStreamService(repo)
	controller := controllers.NewStreamController(service)

	r.GET("/stream/:videoKey", controller.StreamVideo)

	r.Run(":8080")
}
