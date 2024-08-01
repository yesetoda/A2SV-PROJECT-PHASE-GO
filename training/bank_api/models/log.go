package models

type Log struct {
	SenderId   int `json:"sender"`
	RecieverId int `json:"reciever"`
	Amount   float64 `json:"amount"`
}
