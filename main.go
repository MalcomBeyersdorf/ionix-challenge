package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"vaccine-api/database"
	"vaccine-api/handlers"
	"vaccine-api/middlewares"
	"vaccine-api/repositories"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer db.Close()

	// Repos
	userRepo := repositories.NewUserRepository(db)
	drugRepo := repositories.NewDrugRepository(db)
	vaccinationRepo := repositories.NewVaccinationRepository(db)

	// Handlers
	loginHandler := handlers.MakeLoginHandler(userRepo)
	signupHandler := handlers.MakeSignupHandler(userRepo)

	createDrugHandler := handlers.MakeCreateDrugHandler(drugRepo)
	updateDrugHandler := handlers.MakeUpdateDrugHandler(drugRepo)
	getAllDrugsHandler := handlers.MakeGetAllDrugsHandler(drugRepo)
	deleteDrugHandler := handlers.MakeDeleteDrugHandler(drugRepo)

	createVaccinationHandler := handlers.MakeCreateVaccinationHandler(vaccinationRepo, drugRepo)
	updateVaccinationHandler := handlers.MakeUpdateVaccinationHandler(vaccinationRepo)
	getAllVaccinationsHandler := handlers.MakeGetAllVaccinationsHandler(vaccinationRepo)
	deleteVaccinationHandler := handlers.MakeDeleteVaccinationHandler(vaccinationRepo)

	// Routes
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", signupHandler)
	http.Handle("/drugs", middlewares.JWTAuthenticationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createDrugHandler(w, r)
		case http.MethodGet:
			getAllDrugsHandler(w, r)
		}
	})))
	http.Handle("/drugs/", middlewares.JWTAuthenticationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateDrugHandler(w, r)
		case http.MethodDelete:
			deleteDrugHandler(w, r)
		}
	})))
	http.Handle("/vaccination", middlewares.JWTAuthenticationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createVaccinationHandler(w, r)
		case http.MethodGet:
			getAllVaccinationsHandler(w, r)
		}
	})))
	http.Handle("/vaccination/", middlewares.JWTAuthenticationMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateVaccinationHandler(w, r)
		case http.MethodDelete:
			deleteVaccinationHandler(w, r)
		}
	})))

	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	log.Printf("Servidor init on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
