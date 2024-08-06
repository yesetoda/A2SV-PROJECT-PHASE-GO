package Controllers

import (
	"bufio"
	Mongo_Database "example/MongoDB_TaskManager/DataBase"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	db         = Mongo_Database.DB{}
	client, _  = db.ConnectToDB()
	collection = client.Database("TaskDB").Collection("Task")
	reader     = bufio.NewReader(os.Stdin)
)

func HomePage(c *gin.Context) {
	content := []string{
		`this is the home for the task manager api`,
		"List of available commands:",
		"___1,GET _____ localhost:8080/ _______________ this shows the home page",
		"___2,GET _____ localhost:8080/tasks __________ this shows all the registered tasks",
		"___3,GET _____ localhost:8080/task/x _________ this shows the task with id == x ",
		"___4,GET _____ localhost:8080/task/filter ____ this will filter the tasks with specific ",
		"___5,POST ____ localhost:8080/tasks __________ add anew task",
		"___6,PATCH	 ___ localhost:8080/tasks/x _______ update the task with id == x",
		"___7,DELETE ___ localhost:8080/tasks/x _______ remove the task with id == x",
	}

	c.IndentedJSON(http.StatusOK, content)
}

func GetAllTasks(c *gin.Context) {
	result := db.ListAlltasks(*collection)
	c.IndentedJSON(http.StatusOK, result)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	result := db.Gettasks(*collection, intId)
	c.IndentedJSON(http.StatusOK, result)

}

func FilterTask(c *gin.Context) {
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	status := c.Request.FormValue("status")
	fmt.Println("this is the title, description, status", title, description, status)
	result := db.FilterBy(title, description, status, *collection)
	c.IndentedJSON(http.StatusOK, result)
}
func UpdateTaskById(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	status := c.Request.FormValue("status")
	day := c.Request.FormValue("day")
	month := c.Request.FormValue("month")
	year := c.Request.FormValue("year")
	valid, result := db.Updatetask(reader, *collection, intId, title, description, status, day, month, year)
	if valid {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusBadRequest, result)

	}

}

func RemoveTaskById(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	valid, result := db.Removetask(intId, *collection)
	if !valid {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusBadRequest, result)
	}

}

func AddNewTask(c *gin.Context) {
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	status := c.Request.FormValue("status")
	id := c.Request.FormValue("id")
	day := c.Request.FormValue("day")
	month := c.Request.FormValue("month")
	year := c.Request.FormValue("year")
	valid, result := db.RegisterNewtasks(reader, *collection, id, title, description, status, day, month, year)
	if valid {
		c.IndentedJSON(http.StatusOK, result)
	} else {

		c.IndentedJSON(http.StatusBadRequest, result)
	}
}
