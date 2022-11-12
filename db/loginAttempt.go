package db

import (
	"github.com/nehonar/twitteringo/models"
	"golang.org/x/crypto/bcrypt"
)

/*
LoginAttempt check login in MongoDB
*/
func LoginAttempt(email string, password string) (models.User, bool) {
	user, foundUser, _ := CheckExistingEmail(email)

	if !foundUser {
		return user, false
	}

	passwordInDb := []byte(user.Password)
	passwordBytes := []byte(password)

	err := bcrypt.CompareHashAndPassword(passwordInDb, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
