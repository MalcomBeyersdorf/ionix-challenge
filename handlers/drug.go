package handlers

import (
	"net/http"
	"encoding/json"
	"vaccine-api/models"
)

func CreateDrugHandler(w http.ResponseWriter, r *http.Request) {
    var drug models.Drug
    err := json.NewDecoder(r.Body).Decode(&drug)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Insertar la droga en la base de datos

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(drug)
}

func GetAllDrugsHandler(w http.ResponseWriter, r *http.Request) {
    // Obtener todas las drogas de la base de datos

    json.NewEncoder(w).Encode(drugs)
}

func UpdateDrugHandler(w http.ResponseWriter, r *http.Request) {
    // Obtener el ID del URL

    var drug models.Drug
    err := json.NewDecoder(r.Body).Decode(&drug)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Actualizar la droga en la base de datos

    json.NewEncoder(w).Encode(drug)
}

func DeleteDrugHandler(w http.ResponseWriter, r *http.Request) {
    // Obtener el ID del URL

    // Eliminar la droga de la base de datos

    w.WriteHeader(http.StatusOK)
}
