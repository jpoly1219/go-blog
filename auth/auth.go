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
	"github.com/google/uuid"
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
		tokenStruct := generateToken(userIdInt)
		tokens := map[string]string{
			"accessToken":  tokenStruct.AccessToken,
			"refreshToken": tokenStruct.RefreshToken,
		}
		json.NewEncoder(w).Encode(tokens)
	} else {
		fmt.Println("No Match!")
	}
}

func generateToken(userId int) *models.Token {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file.")
	}
	accessSecretKey := os.Getenv("ACCESSSECRETKEY")
	refreshSecretKey := os.Getenv("REFRESHSECRETKEY")

	tokenInfo := &models.Token{}
	tokenInfo.AccessUuid = uuid.NewString()
	tokenInfo.AccessExpire = time.Now().Add(time.Minute * 15).Unix()
	tokenInfo.RefreshUuid = uuid.NewString()
	tokenInfo.RefreshExpire = time.Now().Add(time.Hour * 24 * 7).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["access_uuid"] = tokenInfo.AccessUuid
	accessTokenClaims["user_id"] = userId
	accessTokenClaims["exp"] = tokenInfo.AccessExpire
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenInfo.AccessToken, err = accessToken.SignedString([]byte(accessSecretKey))
	if err != nil {
		panic(err.Error())
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["refresh_uuid"] = uuid.NewString()
	refreshTokenClaims["user_id"] = userId
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenInfo.RefreshToken, err = refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		panic(err.Error())
	}

	return tokenInfo
}
