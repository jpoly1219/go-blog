package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	Content string `json:"Content"`
}

var Posts = make([]Post, 0)

func returnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning all posts...")
	json.NewEncoder(w).Encode(Posts)
}

func returnSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	for _, post := range Posts {
		if post.Id == keys {
			json.NewEncoder(w).Encode(post)
		}
		if post.Id != keys {
			w.Write([]byte(`{"message": "post not found"}`))
			break
		}
	}
}

func createNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	Posts = append(Posts, newPost)
	json.NewEncoder(w).Encode(newPost)
	fmt.Println(Posts)
}

func main() {
	Posts = append(
		Posts,
		Post{Id: "1", Title: "Post1", Author: "Author1", Content: "Content1"},
		Post{Id: "2", Title: "Post2", Author: "Author2", Content: "Content2"},
		Post{Id: "3", Title: "Post3", Author: "Author1", Content: "Content3"},
	)

	r := mux.NewRouter()
	r.HandleFunc("/posts", returnAllPosts)
	r.HandleFunc("/posts/{id}", returnSinglePost)
	r.HandleFunc("/post", createNewPost).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8090", r))
}
