package data

import (
	"example/task_manager_api/models"
	"strconv"
	"strings"
)

type TaskController struct {
	TaskList map[int]models.Task
}

func (t *TaskController) GetAll() map[int]models.Task {
	return t.TaskList
}

func (t *TaskController) GetById(id int) (models.Task, bool) {

	task, tfound := t.TaskList[id]
	if !tfound {
		return task, false
	}
	return task, true

}
func (t *TaskController) Update(updateTask models.Task) (models.Task, bool) {
	t.TaskList[updateTask.ID] = updateTask
	delete(t.TaskList, updateTask.ID)
	t.TaskList[updateTask.ID] = updateTask
	return updateTask,true

}

func (t *TaskController) UpdateField(id int, updateTask models.Task) (models.Task, string) {

	var task models.Task

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
	t.TaskList[id] = task

	delete(t.TaskList, task.ID)
	t.TaskList[task.ID] = task
	return task, "Field Updated successfully"

}
func (t *TaskController) Delete(id int) string {
	delete(t.TaskList, id)
	return "deleted succesfully"

}
func (t *TaskController) Post(task models.Task) string {
	t.TaskList[task.ID] = task
	return "sucessfully Added a new task"

}
