package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jpoly1219/go-blog/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	queryStr := fmt.Sprintf(
		"INSERT INTO users(name, email, username, password) VALUES('%s', '%s', '%s', '%s')",
		user.Name, user.Email, user.Username, user.Password,
	)

	results, err := models.Db.Query(queryStr)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&user.Name, &user.Email, &user.Username, &user.Password)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(user)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(user.Id)

	queryStr := fmt.Sprintf(
		"SELECT * FROM users WHERE email = '%s' AND password = '%s'",
		user.Email, user.Password,
	)

	results, err := models.Db.Query(queryStr)
	if err != nil {
		panic(err.Error())
	}

	type queryOutput struct {
		id, name, email, username, password string
	}

	var qo queryOutput

	for results.Next() {
		err = results.Scan(&qo.id, &qo.name, &qo.email, &qo.username, &qo.password)
		if err != nil {
			panic(err.Error())
		}
	}
	if qo.email == user.Email && qo.password == user.Password {
		fmt.Println("Match!")
	} else {
		fmt.Println("No Match!")
	}

}
