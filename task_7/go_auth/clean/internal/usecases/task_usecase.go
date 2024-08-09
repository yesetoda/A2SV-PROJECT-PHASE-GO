package usecases

import (
	"example/clean/internal/entities"
	"example/clean/internal/interfaces/db"
	"time"
)


type TaskUseCase struct{
	TaskRepo db.TaskRepository
}


func (uc *TaskUseCase) AddTask(id int,title,description,status string,duedate time.Time)(*entities.Task,error){
	Task := &entities.Task{
		Id: id,
		Title: title,
		Description: description,
		Duedate: duedate,
	}
	return uc.TaskRepo.Save(Task)
}
func (uc *TaskUseCase) RemoveTask(id int,title,description,status string,duedate time.Time)(*entities.Task,error){
	Task := &entities.Task{
		Id: id,
		Title: title,
		Description: description,
		Duedate: duedate,
	}
	return uc.TaskRepo.Save(Task)
}
func (uc *TaskUseCase) UpdateTask(id int,title,description,status string,duedate time.Time)(*entities.Task,error){
	Task := &entities.Task{
		Id: id,
		Title: title,
		Description: description,
		Duedate: duedate,
	}
	return uc.TaskRepo.Save(Task)
}
func (uc *TaskUseCase) ViewAllTasks(id int,title,description,status string,duedate time.Time)(*entities.Task,error){
	Task := &entities.Task{
		Id: id,
		Title: title,
		Description: description,
		Duedate: duedate,
	}
	return uc.TaskRepo.Save(Task)
}
func (uc *TaskUseCase) FilterTask(id int,title,description,status string,duedate time.Time)(*entities.Task,error){
	Task := &entities.Task{
		Id: id,
		Title: title,
		Description: description,
		Duedate: duedate,
	}
	return uc.TaskRepo.Save(Task)
}