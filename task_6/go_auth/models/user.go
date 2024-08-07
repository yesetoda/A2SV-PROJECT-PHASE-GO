package models

type User struct {
	UserName string `json:"username"`
	Password string `json:"_"`
	Role     string `json:"role"`
}

type UserCollection struct {
	Users []User
}
