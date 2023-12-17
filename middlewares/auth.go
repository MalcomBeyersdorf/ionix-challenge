package middlewares

import (
    "github.com/dgrijalva/jwt-go"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Método de firma inesperado: %v", token.Header["alg"])
            }

            return []byte("clave"), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            r.Context = context.WithValue(r.Context, "userID", claims["user_id"])
            next(w, r)
        } else {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }
    }
}
