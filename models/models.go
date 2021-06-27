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

type OtherUser struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Token struct {
	AccessToken   string `json:"accesstoken"`
	RefreshToken  string `json:"refreshtoken"`
	AccessUuid    string `json:"accessuuid"`
	RefreshUuid   string `json:"refreshuuid"`
	AccessExpire  int64  `json:"accessexpire"`
	RefreshExpire int64  `json:"refreshexpire"`
}
