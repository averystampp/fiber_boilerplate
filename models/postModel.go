package post

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Author string `json:"author" formdata:"author"`
	Title  string `json:"title" formdata:"title"`
	Body   string `json:"body" formdata:"body"`
}

type User struct {
	gorm.Model
	Username string `json:"username" formdata:"username"`
	Password string `json:"password" formdata:"password"`
}
