package helper

import (
	"fmt"
	"go-api/environment"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	environment.ExportEnv()
	keypass := os.Getenv("KEYPASS")

	convertKeypass, err := strconv.Atoi(keypass)
	if err != nil {
		return "failed convert string to integer", err
	}
	// Generate a hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), convertKeypass)
	if err != nil {
		return "error hashpassword", err
	}

	return string(hashedPassword), nil
}

func DecryptPassword(hashedPassword string, passwordInput string) (string) {
	

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordInput))
	if err != nil {
		fmt.Println("Passwords do not match.")
		return "error hashpassword"
	}

	return "password match"
}
