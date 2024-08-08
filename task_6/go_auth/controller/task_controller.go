package controller

import (
	"example/go_auth/models"
	"example/go_auth/services"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	dbs       = services.DB{}
	secretKey = os.Getenv("MySecret")
	client, _ = dbs.ConnetTOCOllection()
)

func HandleFindTask(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)
	tasks := dbs.Gettasks(*TaskCollection, intId)
	if len(tasks) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "no task with such id"})
	} else {
		c.IndentedJSON(http.StatusOK, tasks[0])
	}

}

func HandleFilterTasks(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	status := c.Request.FormValue("status")
	duedate := c.Request.FormValue("duedate")
	tasks := dbs.FilterBy(title, description, status, duedate, *TaskCollection)
	if len(tasks) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "no task with the desired filter"})
	} else {
		c.IndentedJSON(http.StatusOK, tasks)
	}

}
func HandleViewTasks(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	tasks := dbs.ListAlltasks(*TaskCollection)
	if len(tasks) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "there is no task available"})
	} else {

		c.IndentedJSON(http.StatusOK, tasks)
	}
}
func HandleAddTask(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	var task models.Task
	id := c.Request.FormValue("id")
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	due_date := c.Request.FormValue("due_date")
	status := c.Request.FormValue("status")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "binding error", "error2": err})
		return
	}
	task.Id = intId
	task.Title = title
	task.Description = description
	task.DueDate = due_date
	task.Status = status

	accepted, message := dbs.RegisterNewtasks(*TaskCollection, task)
	if !accepted {
		c.IndentedJSON(http.StatusBadRequest, message)
	} else {
		c.IndentedJSON(http.StatusCreated, "task added succesfully")
	}
}
func HandleEditTask(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	id := c.Param("id")
	title := c.Request.FormValue("title")
	description := c.Request.FormValue("description")
	due_date := c.Request.FormValue("due_date")
	status := c.Request.FormValue("status")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	accepted, message := dbs.Updatetask(*TaskCollection, intId, title, description, status, due_date)
	if !accepted {
		c.IndentedJSON(http.StatusBadRequest, message)
	} else {

		c.IndentedJSON(http.StatusOK, "task edited successful")
	}
}
func HandleRemoveTask(c *gin.Context) {
	TaskCollection := client.Database("JWT_Database").Collection("Tasks")
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {

		accepted, message := dbs.Removetask(intId, *TaskCollection)
		if !accepted {
			c.IndentedJSON(http.StatusNotFound, message)
		} else {

			c.IndentedJSON(http.StatusOK, "task removed successful")
		}
	}

}
