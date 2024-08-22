package services

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(raw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func VerifyHash(hash string, raw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
}