package main

import (
	"log"
	"net/http"

	"github.com/jpoly1219/go-blog/controllers"

	"github.com/gorilla/mux"
)

func main() {
	controllers.Posts = append(
		controllers.Posts,
		controllers.Post{Id: "1", Title: "Post1", Author: "Author1", Content: "Content1"},
		controllers.Post{Id: "2", Title: "Post2", Author: "Author2", Content: "Content2"},
		controllers.Post{Id: "3", Title: "Post3", Author: "Author1", Content: "Content3"},
	)

	r := mux.NewRouter()
	r.HandleFunc("/posts", controllers.ReturnAllPosts)
	r.HandleFunc("/posts/{id}", controllers.ReturnSinglePost).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", controllers.DeletePost).Methods(http.MethodDelete)
	r.HandleFunc("/post", controllers.CreateNewPost).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8090", r))
}
