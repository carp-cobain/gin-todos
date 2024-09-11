package main

import (
	"io"
	"log"

	"github.com/carp-cobain/gin-todos/handlers"
	"github.com/carp-cobain/gin-todos/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDbAndMigrate()

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	v1 := r.Group("/todos/api/v1")
	{
		v1.GET("/stories", handlers.ListStories)
		v1.GET("/stories/:id", handlers.GetStory)
		v1.POST("/stories", handlers.CreateStory)
		v1.PATCH("/stories/:id", handlers.UpdateStory)
		v1.DELETE("/stories/:id", handlers.DeleteStory)
	}

	if err := r.Run(); err != nil {
		log.Panicf("unable to start server:  %+v", err)
	}
}
