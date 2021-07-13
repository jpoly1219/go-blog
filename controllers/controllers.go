package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jpoly1219/go-blog/auth"
	"github.com/jpoly1219/go-blog/models"

	"github.com/gorilla/mux"
)

func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning all posts...")
	var posts = make([]models.Post, 0)

	results, err := models.Db.Query("SELECT id, title, content FROM posts ORDER BY id DESC;")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var post models.Post
		err := results.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	nameResults, err := models.Db.Query("SELECT username FROM users INNER JOIN posts ON users.id=posts.author ORDER BY posts.id DESC;")
	if err != nil {
		panic(err.Error())
	}
	index := 0
	for nameResults.Next() {
		if index < len(posts) {
			err := nameResults.Scan(&posts[index].Author)
			if err != nil {
				panic(err.Error())
			}
			index++
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(posts)
}

func ReturnBatchPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keys := vars["id"]
	fmt.Printf("returning batch %s\n", keys)

	keysInt, _ := strconv.Atoi(keys)
	results, err := models.Db.Query(
		"SELECT id, title, content FROM posts ORDER BY id DESC LIMIT ?, 5;",
		keysInt*5,
	)
	if err != nil {
		panic(err.Error())
	}

	var posts = make([]models.Post, 0)
	for results.Next() {
		var post models.Post
		err := results.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	nameResults, err := models.Db.Query("SELECT username FROM users INNER JOIN posts ON users.id=posts.author ORDER BY posts.id DESC;")
	if err != nil {
		panic(err.Error())
	}
	index := 0
	for nameResults.Next() {
		if index < len(posts) {
			err := nameResults.Scan(&posts[index].Author)
			if err != nil {
				panic(err.Error())
			}
			index++
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(posts)
}

func ReturnSinglePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("returning one post...")
	vars := mux.Vars(r)
	keys := vars["id"]

	results, err := models.Db.Query(
		"SELECT id, title, content FROM posts WHERE id=?;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}
	var post models.Post
	for results.Next() {
		err = results.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			panic(err.Error())
		}
	}

	nameResults, err := models.Db.Query(
		"SELECT username FROM users INNER JOIN posts ON users.id=posts.author WHERE posts.id=?;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}
	for nameResults.Next() {
		err := nameResults.Scan(&post.Author)
		if err != nil {
			panic(err.Error())
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(post)
}

func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	_, err := auth.CheckTokenValidity(r, os.Getenv("ACCESSSECRETKEY"))
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)

	usernameToIdResults, err := models.Db.Query(
		"SELECT id FROM users WHERE username=?;",
		post.Author,
	)
	if err != nil {
		panic(err.Error())
	}
	for usernameToIdResults.Next() {
		err = usernameToIdResults.Scan(&post.Author)
		if err != nil {
			panic(err.Error())
		}
	}

	results, err := models.Db.Query(
		"INSERT INTO posts(title, author, content) VALUES (?, ?, ?);",
		post.Title, post.Author, post.Content,
	)
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
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	_, err := auth.CheckTokenValidity(r, os.Getenv("ACCESSSECRETKEY"))
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	keys := vars["id"]

	var updatedPost models.Post
	json.NewDecoder(r.Body).Decode(&updatedPost)

	results, err := models.Db.Query(
		"UPDATE posts SET title=?, content=? WHERE id=?;",
		updatedPost.Title, updatedPost.Content, keys,
	)
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
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	_, err := auth.CheckTokenValidity(r, os.Getenv("ACCESSSECRETKEY"))
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	keys := vars["id"]

	_, err = models.Db.Query(
		"DELETE FROM posts WHERE id=?;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}
}

func ReturnUserData(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	fmt.Println("returning user data...")
	vars := mux.Vars(r)
	keys := vars["username"]

	results, err := models.Db.Query(
		"SELECT id, name, email, username FROM users WHERE username=?;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}

	var user models.OtherUser
	for results.Next() {
		err = results.Scan(&user.Id, &user.Name, &user.Email, &user.Username)
		if err != nil {
			panic(err.Error())
		}
	}

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(user)
}

func ReturnUserPosts(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	fmt.Println("returning this user's post...")
	vars := mux.Vars(r)
	keys := vars["userid"]

	results, err := models.Db.Query(
		"SELECT * FROM posts WHERE author=? ORDER BY id DESC;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}

	var post models.Post
	var posts []models.Post
	for results.Next() {
		err = results.Scan(&post.Id, &post.Title, &post.Author, &post.Content)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	nameResults, err := models.Db.Query(
		"SELECT username FROM users INNER JOIN posts ON users.id=posts.author WHERE author=?;",
		keys,
	)
	if err != nil {
		panic(err.Error())
	}
	index := 0
	for nameResults.Next() {
		if index < len(posts) {
			err := nameResults.Scan(&posts[index].Author)
			if err != nil {
				panic(err.Error())
			}
			index++
		}
	}

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(posts)
}

func EditUserData(w http.ResponseWriter, r *http.Request) {
	auth.HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	_, err := auth.CheckTokenValidity(r, os.Getenv("ACCESSSECRETKEY"))
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}

	vars := mux.Vars(r)
	keys := vars["username"]

	var updatedUser models.User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	results, err := models.Db.Query(
		"UPDATE users SET name=?, email=?, password=? WHERE id=?;",
		updatedUser.Name, updatedUser.Email, updatedUser.Password, keys,
	)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Email, &updatedUser.Username, &updatedUser.Password)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(updatedUser)
}
