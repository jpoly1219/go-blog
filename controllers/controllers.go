package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jpoly1219/go-blog/models"

	"github.com/gorilla/mux"
)

var Posts = make([]models.Post, 0)

func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning all posts...")
	json.NewEncoder(w).Encode(Posts)
}

func ReturnSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	for _, post := range Posts {
		if post.Id == keys {
			json.NewEncoder(w).Encode(post)
		}
	}
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post
	json.NewDecoder(r.Body).Decode(&newPost)
	Posts = append(Posts, newPost)
	json.NewEncoder(w).Encode(newPost)
	fmt.Println(Posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	for i, post := range Posts {
		if post.Id == keys {
			fmt.Println(keys)
			Posts = append(Posts[:(i)], Posts[i+1:]...)
			json.NewEncoder(w).Encode(post)
		}
	}
	fmt.Println(Posts)
}
