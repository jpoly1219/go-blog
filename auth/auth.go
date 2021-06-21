package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jpoly1219/go-blog/models"
)

func HandleCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://jpoly1219devbox.xyz:5000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	queryStr := fmt.Sprintf(
		"INSERT INTO users(name, email, username, password) VALUES('%s', '%s', '%s', '%s')",
		user.Name, user.Email, user.Username, user.Password,
	)
	results, err := models.Db.Query(queryStr)
	if err != nil {
		fmt.Println("query failed")
	}

	for results.Next() {
		err = results.Scan(&user.Name, &user.Email, &user.Username, &user.Password)
		if err != nil {
			fmt.Println("scan failed; check the number of values in destination and the number of columns")
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(user)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	fmt.Println("login json received...")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	queryStr := fmt.Sprintf(
		"SELECT * FROM users WHERE email = '%s' AND password = '%s'",
		user.Email, user.Password,
	)
	results, err := models.Db.Query(queryStr)
	if err != nil {
		fmt.Println("query failed")
	}

	type queryOutput struct {
		id, name, email, username, password string
	}
	var qo queryOutput
	for results.Next() {
		err = results.Scan(&qo.id, &qo.name, &qo.email, &qo.username, &qo.password)
		if err != nil {
			fmt.Println("scan failed; check the number of values in destination and the number of columns")
		}
	}

	if qo.email == user.Email && qo.password == user.Password {
		fmt.Println("Match!")

		userIdInt, _ := strconv.Atoi(qo.id)
		userName := qo.name
		tokenStruct, err := generateToken(userName, userIdInt)
		if err != nil {
			fmt.Println("failed to generate token")
		}
		accessToken := tokenStruct.AccessToken
		refreshToken := tokenStruct.RefreshToken
		cookie := http.Cookie{
			HttpOnly: true,
			Name:     "refreshToken",
			Value:    refreshToken,
			Domain:   "jpoly1219devbox.xyz",
			Path:     "/auth/",
		}

		responseJSON := map[string]string{
			"accessToken": accessToken,
		}

		http.SetCookie(w, &cookie)
		json.NewEncoder(w).Encode(responseJSON)
	} else {
		fmt.Println("No Match!")
		json.NewEncoder(w).Encode("This user does not exist.")
	}
}

func generateToken(userName interface{}, userId interface{}) (*models.Token, error) {
	var err error

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
	accessTokenClaims["user_name"] = userName
	accessTokenClaims["user_id"] = userId
	accessTokenClaims["exp"] = tokenInfo.AccessExpire
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenInfo.AccessToken, err = accessToken.SignedString([]byte(accessSecretKey))
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["refresh_uuid"] = uuid.NewString()
	refreshTokenClaims["user_name"] = userName
	refreshTokenClaims["user_id"] = userId
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenInfo.RefreshToken, err = refreshToken.SignedString([]byte(refreshSecretKey))
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return tokenInfo, nil
}

func ExtractToken(r *http.Request) string {
	tokenHeaderStr := r.Header.Get("Authorization")
	strSlice := strings.Split(tokenHeaderStr, " ")
	var tokenStr string
	if len(strSlice) == 2 {
		tokenStr = strSlice[1]
	}

	fmt.Println(tokenStr)
	return tokenStr
}

func VerifyToken(r *http.Request, secretKey string) (*jwt.Token, error) {
	tokenStr := ExtractToken(r)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("Failed to verify token.")
		return nil, err
	}

	return token, nil
}

func CheckTokenValidity(r *http.Request, secretKey string) (*jwt.Token, error) {
	token, err := VerifyToken(r, secretKey)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		fmt.Println("Invalid token.")
		return nil, err
	}
	return token, nil
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	HandleCors(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	fmt.Println("refresh request received...")

	// get refreshToken from cookie
	var refreshToken string
	c, err := r.Cookie("refreshToken")
	if err != nil {
		w.Write([]byte("error in reading cookie : " + err.Error() + "\n"))
	} else {
		refreshToken = c.Value
	}

	// verify refreshToken and check its validity
	refreshSecretKey := os.Getenv("REFRESHSECRETKEY")
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(refreshSecretKey), nil
	})
	if err != nil {
		fmt.Println("Failed to verify token.")
		return
	}
	if !token.Valid {
		fmt.Println("Invalid token.")
		return
	}

	var userName interface{}
	var userId interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userName = claims["user_name"]
		userId = claims["user_id"]
	} else {
		fmt.Println(err)
	}

	// generate new access and refresh tokens
	tokenStruct, err := generateToken(userName, userId)
	if err != nil {
		fmt.Println("failed to generate token")
	}
	accessToken := tokenStruct.AccessToken
	refreshToken = tokenStruct.RefreshToken
	cookie := http.Cookie{
		HttpOnly: true,
		Name:     "refreshToken",
		Value:    refreshToken,
		Domain:   "jpoly1219devbox.xyz",
		Path:     "/auth/",
	}

	responseJSON := map[string]string{
		"accessToken": accessToken,
	}

	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode(responseJSON)
}
