package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nehonar/twitteringo/db"
	"github.com/nehonar/twitteringo/models"
)

/*
SignUpUser is a function to check if user exists and sign up if is a new user
*/
func SignUpUser(w http.ResponseWriter, r *http.Request) {
	log.Print("Signup new user")
	var modelUser models.User

	err := json.NewDecoder(r.Body).Decode(&modelUser)
	if err != nil {
		log.Fatal("Error: ", err)
		http.Error(w, "Wrong data sended", 400)
		return
	}

	if len(modelUser.Email) == models.InvalidLengthUserEmail {
		http.Error(w, "Required valid email", 400)
		return
	}

	if len(modelUser.Password) < models.InvalidLengthUserPass {
		http.Error(w, "Required minimum 6 characters in password", 400)
		return
	}

	_, existingEmail, _ := db.CheckExistingEmail(modelUser.Email)
	if existingEmail {
		http.Error(w, "Email in use", 400)
		return
	}

	_, status, err := db.InsertUser(modelUser)
	if err != nil {
		log.Fatal("Error: ", err)
		http.Error(w, "Error in MongoDB when try to insert new user", 400)
		return
	}

	if !status {
		http.Error(w, "Don't save user in MongoDB", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
