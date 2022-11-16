package helper

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func GetPwd(password string) []byte {
	// Prompt the user to enter a password
	// Variable to store the users input
	// Read the users input

	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(password)
}

func HashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// validate strong password
func IsStrongPassword(password string) bool {
	// at least one number, one lowercase and one uppercase letter
	// at least eight characters
	var hasNumber = regexp.MustCompile(`[0-9]+`).MatchString
	var hasLowerChar = regexp.MustCompile(`[a-z]+`).MatchString
	var hasUpperChar = regexp.MustCompile(`[A-Z]+`).MatchString
	var hasMinEightChar = regexp.MustCompile(`.{8,}`).MatchString

	return hasNumber(password) && hasLowerChar(password) && hasUpperChar(password) && hasMinEightChar(password)
}
