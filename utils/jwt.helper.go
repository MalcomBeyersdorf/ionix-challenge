package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID int) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
