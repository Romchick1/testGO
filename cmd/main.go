package main

import (
	"log"
	"net/http"

	"github.com/Romchick1/testGO/internal/handlers"
	"github.com/Romchick1/testGO/internal/repository"
	"github.com/gorilla/mux"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal((err))
	}
	defer repository.DB.Close()

	repo := repository.NewRepository(db)
	productHandler := handlers.NewProductHandler(repo)
	measureHandler := handlers.NewMeasureHandler(repo)

	r := mux.NewRouter()

	// Product routes
	r.HandleFunc("/product/", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/product/", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// Measure routes
	r.HandleFunc("/measure/", measureHandler.GetMeasures).Methods("GET")
	r.HandleFunc("/measure/{id}", measureHandler.GetMeasure).Methods("GET")
	r.HandleFunc("/measure/", measureHandler.CreateMeasure).Methods("POST")
	r.HandleFunc("/measure/{id}", measureHandler.UpdateMeasure).Methods("PUT")
	r.HandleFunc("/measure/{id}", measureHandler.DeleteMeasure).Methods("DELETE")

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
