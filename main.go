package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jpoly1219/go-blog/auth"
	"github.com/jpoly1219/go-blog/controllers"
	"github.com/jpoly1219/go-blog/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file.")
	}

	return os.Getenv(key)
}

func main() {
	dbUsername := getEnvVar("DBUSERNAME")
	dbPassword := getEnvVar("DBPASSWORD")
	dbSource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/goblog", dbUsername, dbPassword)

	var err error
	models.Db, err = sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatalln("Failed to connect to database.")
	}

	defer models.Db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/posts", controllers.ReturnAllPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts/batch/{id}", controllers.ReturnBatchPosts).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", controllers.ReturnSinglePost).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods(http.MethodOptions, http.MethodPut)
	r.HandleFunc("/posts/{id}", controllers.DeletePost).Methods(http.MethodOptions, http.MethodDelete)
	r.HandleFunc("/post", controllers.CreateNewPost).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/{username}", controllers.ReturnUserData).Methods(http.MethodGet)
	r.HandleFunc("/{userid}/posts", controllers.ReturnUserPosts).Methods(http.MethodGet)

	authR := r.PathPrefix("/auth").Subrouter()
	authR.HandleFunc("/signup", auth.SignUp).Methods(http.MethodOptions, http.MethodPost)
	authR.HandleFunc("/login", auth.LogIn).Methods(http.MethodOptions, http.MethodPost)
	authR.HandleFunc("/refresh", auth.Refresh).Methods(http.MethodOptions, http.MethodPost)

	log.Fatal(http.ListenAndServe(":8090", r))
}
