package models

type Task struct {
	Id          int    `bson:"id"`
	Title       string `bson:"title"`
	Description string `bson:"description"`
	DueDate     string `bson:"due_date"`
	Status      string `bson:"status"`
}
