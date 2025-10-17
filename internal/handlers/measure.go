package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Romchick1/testGO/internal/models"
	"github.com/Romchick1/testGO/internal/repository"
	"github.com/gorilla/mux"
)

type MeasureHandler struct {
	repo *repository.Repository
}

func NewMeasureHandler(repo *repository.Repository) *MeasureHandler {
	return &MeasureHandler{repo: repo}
}

func (h *MeasureHandler) GetMeasures(w http.ResponseWriter, r *http.Request) {
	measures, err := h.repo.GetAllMeasures()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(measures)
}

func (h *MeasureHandler) GetMeasure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	measure, err := h.repo.GetMeasureByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(measure)
}

func (h *MeasureHandler) CreateMeasure(w http.ResponseWriter, r *http.Request) {
	var m models.Measure
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	id, err := h.repo.CreateMeasure(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *MeasureHandler) UpdateMeasure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var m models.Measure
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.repo.UpdateMeasure(id, m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *MeasureHandler) DeleteMeasure(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	if err := h.repo.DeleteMeasure(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
