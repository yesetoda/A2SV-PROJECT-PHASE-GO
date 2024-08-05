package Controllers

import (
	Mongo_Database "example/MongoDB_TaskManager/DataBase"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db          = Mongo_Database.DB{}
	client, _ = db.ConnectToDB()
	collection  = client.Database("TaskDB").Collection("Task")
)

func HomePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the home for the task manager api")
}

func GetAllTasks(c *gin.Context) {
	db.Gettasks(*collection)
	c.IndentedJSON(http.StatusOK, "this is the GetAllTasks for the task manager api")
}

func GetTaskById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the hGetTaskByIdome for the task manager api")
}

func FilterTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the FilterTask for the task manager api")
}
func UpdateTaskById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the UpdateTaskById for the task manager api")
}

func RemoveTaskById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is the RemoveTaskById for the task manager api")
}
