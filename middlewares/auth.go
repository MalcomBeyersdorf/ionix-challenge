package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JWTAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			http.Error(w, "Missing auth token", http.StatusForbidden)
			return
		}
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			http.Error(w, "Invalid/Malformed auth token", http.StatusForbidden)
			return
		}
		tokenPart := splitted[1]
		tk := &Claims{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			http.Error(w, "Malformed authentication token", http.StatusForbidden)
			return
		}
		if !token.Valid {
			http.Error(w, "Token is not valid", http.StatusForbidden)
			return
		}
		fmt.Sprintf("User %d", tk.UserID)
		next.ServeHTTP(w, r)
	})
}

type Claims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}
