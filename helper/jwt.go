package helper

import (
	"go-api/environment"
	"go-api/intface"
	"os"
	"github.com/dgrijalva/jwt-go"
)

func JwtSign(p *intface.I_LoginResult) (string, error) {
	payload := jwt.MapClaims{
		"email": p.Email,
		"id":    p.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	environment.ExportEnv()
	keyJwt := os.Getenv("KEYJWT")

	// Sign the token with a secret key
	secretKey := []byte(keyJwt)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "Error signing token:", err
	}

	return signedToken, nil

}
