package helper

import (
	"errors"
	"go-api/environment"
	"go-api/intface"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func JwtSign(p *intface.CheckAccount) (string, error) {
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


func DecryptJWT(tokenString string, secretKey []byte) (*intface.JwtClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &intface.JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		// Return the secret key for validation
		return secretKey, nil
	})

	// Check for errors during parsing
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*intface.JwtClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}