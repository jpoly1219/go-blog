package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jpoly1219/go-blog/auth"
	"github.com/jpoly1219/go-blog/models"

	"github.com/gorilla/mux"
)

func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning all posts...")
	var posts = make([]models.Post, 0)

	results, err := models.Db.Query("SELECT * FROM posts;")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(posts)
}

func ReturnSinglePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]

	results, err := models.Db.Query(fmt.Sprintf("SELECT * FROM posts WHERE id = %s;", keys))
	if err != nil {
		panic(err.Error())
	}

	var post models.Post
	for results.Next() {
		err = results.Scan(&post.Id, &post.Title, &post.Author, &post.Content)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(post)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	err := auth.CheckTokenValidity(r)
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	queryStr := fmt.Sprintf(
		"INSERT INTO posts(title, author, content) VALUES ('%s', '%s', '%s');",
		post.Title, post.Author, post.Content,
	)

	results, err := models.Db.Query(queryStr)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&post.Title, &post.Author, &post.Content)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	err := auth.CheckTokenValidity(r)
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	keys := vars["id"]

	var updatedPost models.Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	queryStr := fmt.Sprintf(
		"UPDATE posts SET title = '%s', content = '%s' WHERE id = %s",
		updatedPost.Title, updatedPost.Content, keys,
	)

	results, err := models.Db.Query(queryStr)
	if err != nil {
		panic(err.Error())
	}

	var post models.Post
	for results.Next() {
		err = results.Scan(&post.Id, &post.Title, &post.Author, &post.Content)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	err := auth.CheckTokenValidity(r)
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	keys := vars["id"]

	_, err = models.Db.Query(fmt.Sprintf("DELETE FROM posts WHERE id = %s", keys))
	if err != nil {
		panic(err.Error())
	}
}
