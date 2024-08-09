package entities

import "time"


type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duedate     time.Time `json:"duedate"`
	Status      string `json:"status"`
}
