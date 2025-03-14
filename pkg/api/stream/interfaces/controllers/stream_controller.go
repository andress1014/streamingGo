package controllers

import (
	"context"
	"goStreaming/pkg/api/stream/interfaces/services"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StreamController struct {
	Service services.StreamService
}

func NewStreamController(service services.StreamService) *StreamController {
	return &StreamController{Service: service}
}

func (c *StreamController) StreamVideo(ctx *gin.Context) {
	videoKey := ctx.Param("videoKey")
	stream, err := c.Service.GetStream(context.Background(), videoKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stream.Close()

	data, err := io.ReadAll(stream)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "video/mp4", data)
}
