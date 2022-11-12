package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nehonar/twitteringo/models"
)

/*
GenerateJWT create encrypt token with JWT
*/
func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("Nehonar_challenge_reply_twitter")

	payload := jwt.MapClaims{
		"email":       user.Email,
		"name":        user.Name,
		"lastname":    user.Lastname,
		"dateOfBirth": user.DateOfBirth,
		"biography":   user.Biography,
		"location":    user.Location,
		"webSite":     user.WebSite,
		"_id":         user.ID.Hex(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
