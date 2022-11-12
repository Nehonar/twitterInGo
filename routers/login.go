package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nehonar/twitteringo/db"
	"github.com/nehonar/twitteringo/jwt"
	"github.com/nehonar/twitteringo/models"
)

/*
Login realize login
*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid user or password"+err.Error(), 400)
		return
	}

	if len(user.Email) == models.InvalidLengthUserEmail {
		http.Error(w, "Email user is required", 400)
		return
	}

	document, exist := db.LoginAttempt(user.Email, user.Password)
	if !exist {
		http.Error(w, "Invalid user or password", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Creating JWT attempt was wrong", 400)
	}

	resposne := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resposne)

	//Save expiration date in cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
