package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jpoly1219/go-blog/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

}

func LogIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	queryStr := fmt.Sprintf(
		"SELECT * FROM users WHERE email = %s AND password = %s",
		user.Email, user.Password,
	)

	results, err := models.Db.Query(queryStr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(results)

}
