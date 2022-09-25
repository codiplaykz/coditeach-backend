package models

type User struct {
	Id       uint   `json:"id"`
	Login    string `json:"login"`
	Role     string `json:"role"`
	Password []byte `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
}
