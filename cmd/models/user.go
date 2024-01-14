package models

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewUser(id int64, name string) User {
	return User{Id: id, Name: name}
}
