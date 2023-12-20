package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vaccine-api/models"
	"vaccine-api/repositories"
)

// Handler para POST /drugs
func MakeCreateDrugHandler(drugRepo *repositories.DrugRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var drug models.Drug
		err := json.NewDecoder(r.Body).Decode(&drug)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		err = drugRepo.CreateDrug(drug)
		if err != nil {
			http.Error(w, "Error creating drug", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// Handler para PUT /drugs/:id
func MakeUpdateDrugHandler(drugRepo *repositories.DrugRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))
		if err != nil {
			http.Error(w, "Invalid drug ID", http.StatusBadRequest)
			return
		}

		var drug models.Drug
		err = json.NewDecoder(r.Body).Decode(&drug)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		err = drugRepo.UpdateDrug(id, drug)
		if err != nil {
			http.Error(w, "Error updating drug", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Handler para GET /drugs
func MakeGetAllDrugsHandler(drugRepo *repositories.DrugRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		drugs, err := drugRepo.GetAllDrugs()
		if err != nil {
			http.Error(w, "Error retrieving drugs", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(drugs)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	}
}

// Handler para DELETE /drugs/:id
func MakeDeleteDrugHandler(drugRepo *repositories.DrugRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get(":id"))
		if err != nil {
			http.Error(w, "Invalid drug ID", http.StatusBadRequest)
			return
		}

		err = drugRepo.DeleteDrug(id)
		if err != nil {
			http.Error(w, "Error deleting drug", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
