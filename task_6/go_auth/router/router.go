package router

import (
	"example/go_auth/controller"
	"example/go_auth/middleware"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()

	// Route for generating tokens

	r.POST("/login", controller.HandleLogin)
	r.POST("/signup", controller.HandleSignUp)

	// Middleware to check JWT on every request
	// TODO: any thing below this is treated as a protected route
	r.Use(middleware.AuthMiddleware())

	// Protected routes
	r.GET("/tasks", middleware.UserMiddleware(), controller.HandleViewTasks)
	r.PATCH("/promote/:username", middleware.AdminMiddleware(), controller.HandlePromote)
	r.POST("/task", middleware.AdminMiddleware(), controller.HandleAddTask)
	r.PATCH("/task/:id", middleware.AdminMiddleware(), controller.HandleEditTask)
	r.DELETE("/task/:id", middleware.AdminMiddleware(), controller.HandleRemoveTask)
	r.GET("/users", middleware.AdminMiddleware(), controller.HandleViewUsers)
	r.Run(":8080")
}
