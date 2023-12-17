package main

import (
	"net/http"
	"vaccine-api/handlers"
)

func main() {

	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/drugs", handlers.GetAllDrugsHandler)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
