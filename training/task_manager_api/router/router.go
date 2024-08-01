package router

import (
	"example/task_manager_api/controllers"

	"github.com/gin-gonic/gin"
)

var ()

func Router(router *gin.Engine, t controllers.TaskController) {
	router.GET("/", func(c *gin.Context) { c.IndentedJSON(200, "this is the task manager api") })
	router.GET("/tasks", t.HandleGetAll)
	router.GET("/tasks/:id", t.HandleGetById)
	router.PUT("/tasks/:id", t.HandleUpdate)
	router.PATCH("/tasks/:id", t.HandleUpdateField)
	router.DELETE("/tasks/:id", t.HandleDelete)
	router.POST("/tasks", t.HandlePost)
}
