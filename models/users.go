package models

import "gorm/db"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Esta estructura de usuarios es una lista [] de usuario. Cada usuario tiene la estructura de User
type Users []User

func MigrarUser() {
	db.Database.AutoMigrate(User{})
}
