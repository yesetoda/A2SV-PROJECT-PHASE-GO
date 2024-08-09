package router

import (
	"example/task_manager_api/controllers"

	"github.com/gin-gonic/gin"
)

var ()

func Router(router *gin.Engine) {
	router.GET("/", controllers.HandleLandingPage)
	router.GET("/tasks",controllers.HandleGetAll)
	router.GET("/tasks/:id",controllers.HandleGetById)
	router.PUT("/tasks/:id",controllers.HandleUpdate)
	router.PATCH("/tasks/:id",controllers.HandleUpdateField)
	router.DELETE("/tasks/:id",controllers.HandleDelete)
	router.POST("/tasks",controllers.HandlePost)
}
