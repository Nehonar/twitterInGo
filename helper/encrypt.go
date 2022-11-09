package helper

import "golang.org/x/crypto/bcrypt"

/*
security is number of laps to encrypt
*/
var security = 8

/*
Encrypt to encrypt something
*/
func Encrypt(dataToEncrypt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(dataToEncrypt), security)
	return string(bytes), err
}
