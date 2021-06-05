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
	var posts = make([]models.Post, 0)

	results, err := models.Db.Query("SELECT * FROM posts")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var post models.Post
		err := results.Scan(&post.Id, &post.Title, &post.Author, &post.Content)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func ReturnSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	var posts = make([]models.Post, 0)

	for _, post := range posts {
		if post.Id == keys {
			json.NewEncoder(w).Encode(post)
		}
	}
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	var newPost models.Post
	json.NewDecoder(r.Body).Decode(&newPost)

	var posts = make([]models.Post, 0)

	posts = append(posts, newPost)
	json.NewEncoder(w).Encode(newPost)
	fmt.Println(posts)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	var posts = make([]models.Post, 0)

	for i, post := range posts {
		if post.Id == keys {
			fmt.Println(keys)
			posts = append(posts[:(i)], posts[i+1:]...)
			json.NewEncoder(w).Encode(post)
		}
	}
	fmt.Println(posts)
}
