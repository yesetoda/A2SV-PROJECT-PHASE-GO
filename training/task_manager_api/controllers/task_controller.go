package controllers

import (
	"example/task_manager_api/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskList map[int]models.Task
}

func (t *TaskController) HandleLandingPage(c *gin.Context) {
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
func (t *TaskController) HandleGetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, t.TaskList)
}

func (t *TaskController) HandleGetById(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	task, tfound := t.TaskList[int_id]
	if !tfound {
		c.IndentedJSON(404, "not task with such id")
		return
	}
	c.IndentedJSON(http.StatusOK, task)

}
func (t *TaskController) HandleUpdate(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	task, tfound := t.TaskList[int_id]
	if !tfound {
		c.IndentedJSON(404, "not task with such id")
		return
	}

	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.Description = updateTask.Description
	task.Title = updateTask.Title
	task.ID = updateTask.ID
	task.DueDate = updateTask.DueDate
	task.Status = updateTask.Status
	t.TaskList[int_id] = task
	delete(t.TaskList, int_id)
	x, _ := strconv.Atoi(task.ID)
	t.TaskList[x] = task
	c.IndentedJSON(http.StatusOK, "Updated successfully")

}

func (t *TaskController) HandleUpdateField(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	task, tfound := t.TaskList[int_id]
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
	if len(strings.Trim(updateTask.ID, " ")) > 0 {
		task.ID = updateTask.ID
	}
	if len(strings.Trim(updateTask.DueDate.GoString(), " ")) > 0 {

		task.DueDate = updateTask.DueDate
	}
	if len(strings.Trim(updateTask.Status, " ")) > 0 {

		task.Status = updateTask.Status
	}
	t.TaskList[int_id] = task

	x, _ := strconv.Atoi(task.ID)
	delete(t.TaskList, int_id)
	t.TaskList[x] = task
	c.IndentedJSON(http.StatusOK, "Field Updated successfully")

}
func (t *TaskController) HandleDelete(c *gin.Context) {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, "id must be an int")
		return
	}
	_, tfound := t.TaskList[int_id]
	if !tfound {
		c.IndentedJSON(404, "not task with such id")
		return
	}
	delete(t.TaskList, int_id)
	c.IndentedJSON(http.StatusOK, "deleted succesfully")

}
func (t *TaskController) HandlePost(c *gin.Context) {

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		c.IndentedJSON(http.StatusBadRequest, newTask)

		return
	}

	int_id, err := strconv.Atoi(newTask.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "id must be an int")
		return
	}
	if _, found := t.TaskList[int_id]; found {
		c.IndentedJSON(http.StatusBadRequest, "the id is taken")
		return
	}
	t.TaskList[int_id] = newTask
	c.IndentedJSON(http.StatusCreated, "sucessfully Added a new task")

}
