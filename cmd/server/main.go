package main

import (
	"go-task-api/internal/config"
	"go-task-api/internal/handlers"
	"go-task-api/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	db := config.ConnectDatabase()

	if err = db.AutoMigrate(&models.Task{}); err != nil {
		panic(err)
	}

	router := gin.Default()

	taskHandler := handlers.TaskHandler{DB: db}

	router.GET("/v1/tasks", taskHandler.GetTasks)
	router.POST("/v1/tasks", taskHandler.CreateTask)
	router.GET("/v1/tasks/:id", taskHandler.GetTask)
	router.PUT("/v1/tasks/:id", taskHandler.UpdateTask)
	router.DELETE("/v1/tasks/:id", taskHandler.DeleteTask)

	router.Run(":8080")
}
