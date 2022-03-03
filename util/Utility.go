package util

import (
	"database/sql"
	"restaurant/config"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func PasswordEncoder(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func PasswordComparator(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePlain := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}

func CheckConnection() *sql.DB {
	db, errdb := config.GetConnection()
	if errdb != nil {
		log.Error(MISSING_CONNECTION)
		log.Fatal(errdb)
	}
	return db
}
