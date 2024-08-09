// hre you're gonna implement the requests from the usecase
package db

import (
	"database/sql"
	"example/clean/internal/entities"
)

type TaskRepositoryInterface interface{
	AddTask()
	RemoveTask()
	UpdaeTask()
	ViewAllTasks()
	FilterTask()
}

type TaskRepository struct {
	dbConn *sql.DB
}

func (repo *TaskRepository) Save(task *entities.Task) (*entities.Task, error) {
	//TODO:  the logic here
	return task, nil
}
