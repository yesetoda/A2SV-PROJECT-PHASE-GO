package controllers

import (
	"example/task_manager_api/data"
	"example/task_manager_api/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)


func HandleLandingPage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"List Of API's Avalilable": []map[string]map[string]string{
		{"GET /": {
			"description": "home",
			"Body":        "",
			"param":       "",
		}},
		{"GET /tasks": {
			"description": "get all the tasks",
			"Body":        "",
			"param":       "",
		}},
		{"GET /tasks/id": {
			"description": "get specifit task by id",
			"param":       "id",
			"Body":        "",
		}},
		{"PUT /tasks/id": {
			"description": "update all fields of specific task",
			"param":       "id",
			"Body":        "all of title,description,due_date,status",
		}},
		{"PATCH /tasks/id": {
			"description": "home",
			"param":       "id",
			"Body":        "any of title,description,due_date,status",
		}},
		{"DELETE /tasks/id": {
			"description": "home",
			"param":       "id",
			"Body":        "",
		}},
		{"POST /tasks": {
			"description": "home",
			"param":       "",
			"Body":        "all of id,title,description,due_date,status",
		}},
	}})
}
var taskController data.TaskController
func HandleGetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, taskController.GetAll())
}

func HandleGetById(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	task,accepted := taskController.GetById(int_id)
	if !accepted {
		c.IndentedJSON(404, "not task with such id")
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}
func HandleUpdate(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	var updateTask models.Task
	updateTask.Description = c.Request.FormValue("description")
	updateTask.Title = c.Request.FormValue("Title")
	updateTask.DueDate = c.Request.FormValue("due_date")
	updateTask.Status = c.Request.FormValue("status")
	updateTask.ID = int_id
	task,accepted := taskController.Update(updateTask)
	if !accepted{
		c.IndentedJSON(http.StatusConflict, "update not accepted")
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func HandleUpdateField(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	task, tfound := taskController.TaskList[int_id]
	if !tfound {
		c.IndentedJSON(404, "not task with such id")
		return
	}

	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(strings.Trim(updateTask.Description, " ")) > 0 {

		task.Description = updateTask.Description
	}
	if len(strings.Trim(updateTask.Title, " ")) > 0 {

		task.Title = updateTask.Title
	}
	if len(strings.Trim(strconv.Itoa(updateTask.ID), " ")) > 0 {
		task.ID = updateTask.ID
	}
	if len(strings.Trim(updateTask.DueDate, " ")) > 0 {

		task.DueDate = updateTask.DueDate
	}
	if len(strings.Trim(updateTask.Status, " ")) > 0 {

		task.Status = updateTask.Status
	}
	taskController.TaskList[int_id] = task
	delete(taskController.TaskList, int_id)
	taskController.TaskList[int_id] = task
	c.IndentedJSON(http.StatusOK, "Field Updated successfully")

}
func HandleDelete(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	_, tfound := taskController.TaskList[int_id]
	if !tfound {
		c.IndentedJSON(404, "not task with such id")
		return
	}
	_ = taskController.Delete(int_id)
	
	c.IndentedJSON(http.StatusOK, "deleted succesfully")

}
func HandlePost(c *gin.Context) {

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.IndentedJSON(http.StatusBadRequest, newTask)

		return
	}

	
	
	if _, found := taskController.TaskList[newTask.ID]; found {
		c.IndentedJSON(http.StatusBadRequest, "the id is taken")
		return
	}
	_ = taskController.Post(newTask)
	c.IndentedJSON(http.StatusCreated, "sucessfully Added a new task")

}
