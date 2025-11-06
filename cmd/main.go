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
	defer db.Close()

	repo := repository.NewRepository(db)
	measureHandler := handlers.NewMeasureHandler(repo)

	r := mux.NewRouter()

	managerHandler := handlers.NewManagerHandler(repo)
	productHandler := handlers.NewProductHandler(repo)

	auth := managerHandler.AuthMiddleware

	r.HandleFunc("/manager/me", auth(managerHandler.GetMyInfo)).Methods("GET")
	r.HandleFunc("/manager/", auth(managerHandler.GetAllManagers)).Methods("GET")
	r.HandleFunc("/manager/", auth(managerHandler.CreateManager)).Methods("POST")
	r.HandleFunc("/manager/{login}", auth(managerHandler.UpdateManager)).Methods("PUT")
	r.HandleFunc("/manager/{login}", auth(managerHandler.DeleteManager)).Methods("DELETE")

	r.HandleFunc("/product/", auth(productHandler.GetProducts)).Methods("GET")
	r.HandleFunc("/product/{id}", auth(productHandler.GetProducts)).Methods("GET")
	r.HandleFunc("/product/", auth(productHandler.CreateProduct)).Methods("POST")
	r.HandleFunc("/product/{id}", auth(productHandler.UpdateProduct)).Methods("PUT")
	r.HandleFunc("/product/{id}", auth(productHandler.DeleteProduct)).Methods("DELETE")

	r.HandleFunc("/measure/", auth(measureHandler.GetMeasures)).Methods("GET")

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
