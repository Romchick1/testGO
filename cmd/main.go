package main

import (
	"log"
	"net/http"

	"github.com/Romchick1/testGO/internal/handlers"
	"github.com/Romchick1/testGO/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	repository.InitDB()
	defer repository.DB.Close()

	r := mux.NewRouter()

	// Product routes
	r.HandleFunc("/product/", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/product/", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", handlers.DeleteProduct).Methods("DELETE")

	// Measure routes
	r.HandleFunc("/measure/", handlers.GetMeasures).Methods("GET")
	r.HandleFunc("/measure/{id}", handlers.GetMeasure).Methods("GET")
	r.HandleFunc("/measure/", handlers.CreateMeasure).Methods("POST")
	r.HandleFunc("/measure/{id}", handlers.UpdateMeasure).Methods("PUT")
	r.HandleFunc("/measure/{id}", handlers.DeleteMeasure).Methods("DELETE")

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
