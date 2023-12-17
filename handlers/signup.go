package handlers

import (
	"net/http"
	"encoding/json"
	"vaccine-api/models"
	"golang.org/x/crypto/bcrypt"
	"vaccine-api/database"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		http.Error(w, "Error al hashear la contraseña", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	db := database.Connect()
	defer db.Close()

	sqlStatement := `
	INSERT INTO users (name, email, password)
	VALUES ($1, $2, $3)
	RETURNING id`

	var id int
	err = db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		http.Error(w, "Error al insertar usuario en la base de datos", http.StatusInternalServerError)
		return
	}

	user.ID = id

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user) 
}
