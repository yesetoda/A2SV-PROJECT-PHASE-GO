package main

import (
	"example/MongoDB_TaskManager/Controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", Controllers.HomePage)
	router.GET("/tasks", Controllers.GetAllTasks)
	router.GET("/task/:id", Controllers.GetTaskById)
	router.GET("/task/filter", Controllers.FilterTask)
	router.POST("/task", Controllers.AddNewTask)
	router.PATCH("/task/:id", Controllers.UpdateTaskById)
	router.DELETE("/task/:id", Controllers.RemoveTaskById)
	router.Run("localhost:8080")

}
