package models

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"task"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
}

