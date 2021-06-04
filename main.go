package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jpoly1219/go-blog/controllers"
	"github.com/jpoly1219/go-blog/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv(key)
}

func main() {
	controllers.Posts = append(
		controllers.Posts,
		models.Post{Id: "1", Title: "Post1", Author: "Author1", Content: "Content1"},
		models.Post{Id: "2", Title: "Post2", Author: "Author2", Content: "Content2"},
		models.Post{Id: "3", Title: "Post3", Author: "Author1", Content: "Content3"},
	)

	dbUsername := getEnvVar("DBUSERNAME")
	dbPassword := getEnvVar("DBPASSWORD")

	dbSource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/posts", dbUsername, dbPassword)
	fmt.Println(dbSource)

	Db, err := sql.Open("mysql", dbSource)

	if err != nil {
		panic(err.Error())
	}

	defer Db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/posts", controllers.ReturnAllPosts)
	r.HandleFunc("/posts/{id}", controllers.ReturnSinglePost).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", controllers.DeletePost).Methods(http.MethodDelete)
	r.HandleFunc("/post", controllers.CreateNewPost).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8090", r))
}
