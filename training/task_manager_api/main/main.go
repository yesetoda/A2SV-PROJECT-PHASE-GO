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
	// taskId = 0
)

func main() {
	fmt.Println("this is the router")
	Router := gin.Default()
	router.Router(Router, t)
	fmt.Println("the router is running on port 8080")
	Router.Run("localhost:9090")

}

// func f(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	w.Write([]byte("this is the handler"))
// }

// func handleGetAll(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	json.NewEncoder(w).Encode(w)

// }

// func handleGetById(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	w.Write([]byte("this is the handler"))
// }
// func handleUpdate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	w.Write([]byte("this is the handler"))
// }
// func handleDelete(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	w.Write([]byte("this is the handler"))
// }
// func handlePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "application/json")
// 	w.Write([]byte("this is the handler"))
// }
