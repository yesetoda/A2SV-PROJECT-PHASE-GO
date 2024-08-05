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
	router.GET("filter", Controllers.FilterTask)
	router.PATCH("/update/:id", Controllers.UpdateTaskById)
	router.DELETE("/remove/:id", Controllers.RemoveTaskById)
	router.Run("localhost:8080")

}
