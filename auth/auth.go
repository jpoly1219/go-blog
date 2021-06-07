package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
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
		userIdInt, _ := strconv.Atoi(qo.id)
		token := generateToken(userIdInt)
		json.NewEncoder(w).Encode(token)
	} else {
		fmt.Println("No Match!")
	}
}

func generateToken(userId int) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file.")
	}
	jwtSecretKey := os.Getenv("JWTSECRETKEY")

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["user_id"] = userId
	accessTokenClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		panic(err.Error())
	}
	return accessTokenString
}
