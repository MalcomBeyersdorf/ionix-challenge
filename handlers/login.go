package handlers

import (
    "net/http"
	"encoding/json"
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
    "vaccine-api/models"

)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds models.User
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // user, err := GetUserByEmail(creds.Email)

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
        http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
    })

    tokenString, err := token.SignedString([]byte("clave"))
    if err != nil {
        http.Error(w, "Error al generar el token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
