package models

type User struct{
	Username string
	Email string
	Password string
	Token string
}

type UserRequest struct{
	Email string `json:"email"`
	Username string `json:"username"`
}