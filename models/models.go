package models

import (
	"database/sql"
)

var Db *sql.DB

type Post struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
