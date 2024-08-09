package main

import (
	"fmt"

	"example/task_manager_api/router"

	"github.com/gin-gonic/gin"
)
func main() {
	fmt.Println("this is the router")
	Router := gin.Default()
	router.Router(Router)
	fmt.Println("the router is running on port 8080")
	Router.Run("localhost:9090")

}
