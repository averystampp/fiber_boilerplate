package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" formdata:"username"`
	Password string `json:"password" formdata:"password"`
}
