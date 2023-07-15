package util

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(userPassword))

	if err != nil {
		log.Print("E-Mail or Password is incorrect")
		return false
	}
	return true
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err.Error())
	}
	return string(bytes)
}
