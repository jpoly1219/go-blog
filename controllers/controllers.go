package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jpoly1219/go-blog/models"

	"github.com/gorilla/mux"
)

func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning all posts...")
	json.NewEncoder(w).Encode(models.Posts)
}

func ReturnSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	for _, post := range models.Posts {
		if post.Id == keys {
			json.NewEncoder(w).Encode(post)
		}
	}
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post
	json.NewDecoder(r.Body).Decode(&newPost)
	models.Posts = append(models.Posts, newPost)
	json.NewEncoder(w).Encode(newPost)
	fmt.Println(models.Posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	for i, post := range models.Posts {
		if post.Id == keys {
			fmt.Println(keys)
			models.Posts = append(models.Posts[:(i)], models.Posts[i+1:]...)
			json.NewEncoder(w).Encode(post)
		}
	}
	fmt.Println(models.Posts)
}
