package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {
	salt := 8
	password := []byte(p)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hashedPassword)
}

func ComparePass(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
