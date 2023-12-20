package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"vaccine-api/models"
	"vaccine-api/repositories"
)

// Handler para POST /vaccination
func MakeCreateVaccinationHandler(vaccinationRepo *repositories.VaccinationRepository, drugRepo *repositories.DrugRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vaccination models.Vaccination
		err := json.NewDecoder(r.Body).Decode(&vaccination)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		err = vaccinationRepo.CreateVaccination(vaccination, drugRepo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// Handler para PUT /vaccination/:id
func MakeUpdateVaccinationHandler(vaccinationRepo *repositories.VaccinationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid vaccination ID", http.StatusBadRequest)
			return
		}

		var vaccination models.Vaccination
		err = json.NewDecoder(r.Body).Decode(&vaccination)
		if err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		err = vaccinationRepo.UpdateVaccination(id, vaccination)
		if err != nil {
			http.Error(w, "Error updating vaccination", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// Handler para GET /vaccination
func MakeGetAllVaccinationsHandler(vaccinationRepo *repositories.VaccinationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		vaccinations, err := vaccinationRepo.GetAllVaccinations()
		if err != nil {
			http.Error(w, "Error retrieving vaccinations", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(vaccinations)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	}
}

// Handler para DELETE /vaccination/:id
func MakeDeleteVaccinationHandler(vaccinationRepo *repositories.VaccinationRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid vaccination ID", http.StatusBadRequest)
			return
		}

		err = vaccinationRepo.DeleteVaccination(id)
		if err != nil {
			http.Error(w, "Error deleting vaccination", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
