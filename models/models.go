package models

import (
	"database/sql"
)

var Db *sql.DB

type Post struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	Content string `json:"Content"`
}
