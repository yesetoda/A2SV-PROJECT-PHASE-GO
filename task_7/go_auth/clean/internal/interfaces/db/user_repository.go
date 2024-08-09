package db

import (
	"database/sql"
	"example/clean/internal/entities"
)

type UserRepository struct{
	dbConn *sql.DB
}

func (repo *UserRepository) Save(user *entities.User)(*entities.User,error){
	//TODO:  the logic here
	return user,nil
}