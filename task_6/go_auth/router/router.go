package router

import (
	"example/go_auth/controller"
	"example/go_auth/middleware"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()
	r.POST("/login", controller.HandleLogin)
	r.POST("/signup", controller.HandleSignUp)

	tasks := r.Group("/task")
	tasks.Use(middleware.AuthMiddleware())
	{
		tasks.GET("/all", middleware.UserMiddleware(), controller.HandleViewTasks)
		tasks.GET("/:id", middleware.UserMiddleware(), controller.HandleFindTask)
		tasks.GET("/filter", middleware.UserMiddleware(), controller.HandleFilterTasks)
		tasks.POST("/", middleware.AdminMiddleware(), controller.HandleAddTask)
		tasks.PATCH("/:id", middleware.AdminMiddleware(), controller.HandleEditTask)
		tasks.DELETE("/:id", middleware.AdminMiddleware(), controller.HandleRemoveTask)
	}
	users := r.Group("/user")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("/all", middleware.AdminMiddleware(), controller.HandleViewAllUsers)
		users.GET("/:role", middleware.AdminMiddleware(), controller.HandleRoleBasedUsersView)
		users.GET("/u/:username", middleware.AdminMiddleware(), controller.HandleFindUsers)
		users.PATCH("/:username", middleware.AdminMiddleware(), controller.HandlePromote)
	}
	r.Run(":8080")
}
