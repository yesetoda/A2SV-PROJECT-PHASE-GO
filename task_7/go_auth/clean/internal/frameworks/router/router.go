package router

import (
	"example/clean/internal/interfaces/thehttp"

	"github.com/gin-gonic/gin"
)

func NewRoute(UserController thehttp.UserController) *gin.Engine{
	r:= gin.Default()
	r.POST("/signup",UserController.CreateUser)
	r.POST("/signin",UserController.CreateUser)
	return r
}
