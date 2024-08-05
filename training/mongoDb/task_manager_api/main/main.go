package main

import (
	"fmt"

	"example/task_manager_api/controllers"
	"example/task_manager_api/models"
	"example/task_manager_api/router"

	"github.com/gin-gonic/gin"
)

var (
	t = controllers.TaskController{
		TaskList: make(map[int]models.Task),
	}
)

func Disconnect(){

}
func main() {
	
	fmt.Println("this is the router")
	Router := gin.Default()
	router.Router(Router, t)
	fmt.Println("the router is running on port 8080")
	Router.Run("localhost:9090")

}