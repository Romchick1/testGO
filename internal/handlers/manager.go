package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Romchick1/testGO/internal/models"
	"github.com/Romchick1/testGO/internal/repository"
	"github.com/gorilla/mux"
)

type ManagerHandler struct {
	repo *repository.Repository
}

func NewManagerHandler(repo *repository.Repository) *ManagerHandler {
	return &ManagerHandler{repo: repo}
}

func (h *ManagerHandler) GetMyInfo(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	manager, err := h.repo.GetManagerByLogin(login)
	if err != nil {
		http.Error(w, "manager not found", http.StatusNotFound)
		return
	}

	resp := models.ManagerResponse{
		Login:    manager.Login,
		FullName: manager.FullName,
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *ManagerHandler) GetAllManagers(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	if login != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	managers, err := h.repo.GetAllManagers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resp []models.ManagerResponse
	for _, m := range managers {
		resp = append(resp, models.ManagerResponse{Login: m.Login, FullName: m.FullName})
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *ManagerHandler) CreateManager(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	if login != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	var m models.Manager
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreateManager(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *ManagerHandler) UpdateManager(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	if login != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	targetLogin := vars["login"]

	var m models.Manager
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdateManager(targetLogin, m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ManagerHandler) DeleteManager(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	if login != "admin" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	targetLogin := vars["login"]

	if err := h.repo.DeleteManager(targetLogin); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
