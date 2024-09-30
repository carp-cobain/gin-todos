package main

import (
	"log"
	"os"

	"github.com/carp-cobain/gin-todos/database"
	"github.com/carp-cobain/gin-todos/database/repo"
	"github.com/carp-cobain/gin-todos/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	if _, ok := os.LookupEnv("DISABLE_COLOR"); ok {
		gin.DisableConsoleColor()
	}

	readConn, writeConn, err := database.ConnectAndMigrate()
	if err != nil {
		log.Panicf("unable to connnect to db: %+v", err)
	}

	storyRepo := repo.NewStoryRepo(readConn, writeConn)
	taskRepo := repo.NewTaskRepo(readConn, writeConn)

	storyHandler := handler.NewStoryHandler(storyRepo)
	taskHandler := handler.NewTaskHandler(taskRepo)

	r := gin.Default()
	v1 := r.Group("/todos/api/v1")
	{
		// Story routes
		v1.GET("/stories", storyHandler.GetStories)
		v1.GET("/stories/:id", storyHandler.GetStory)
		v1.POST("/stories", storyHandler.CreateStory)
		v1.PATCH("/stories/:id", storyHandler.UpdateStory)
		v1.DELETE("/stories/:id", storyHandler.DeleteStory)
		v1.GET("/stories/:id/tasks", taskHandler.GetTasks)

		// Task routes
		v1.GET("/tasks/:id", taskHandler.GetTask)
		v1.POST("/tasks", taskHandler.CreateTask)
		v1.PATCH("/tasks/:id", taskHandler.UpdateTask)
		v1.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}

	if err := r.Run(); err != nil {
		log.Panicf("unable to start server:  %+v", err)
	}
}
