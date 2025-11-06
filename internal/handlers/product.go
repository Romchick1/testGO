package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Romchick1/testGO/internal/models"
	"github.com/Romchick1/testGO/internal/repository"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	repo *repository.Repository
}

func NewProductHandler(repo *repository.Repository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	manager, err := h.repo.GetManagerByLogin(login)
	if err != nil {
		http.Error(w, "manager not found", http.StatusNotFound)
		return
	}

	products, err := h.repo.GetProductsByManagerID(manager.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	manager, err := h.repo.GetManagerByLogin(login)
	if err != nil {
		http.Error(w, "manager not found", http.StatusNotFound)
		return
	}

	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	p.ManagerID = manager.ID

	id, err := h.repo.CreateProduct(p)
	if err != nil {
		if strings.Contains(err.Error(), "unique_product_name") {
			http.Error(w, "product name must be unique", http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	manager, err := h.repo.GetManagerByLogin(login)
	if err != nil {
		http.Error(w, "manager not found", http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	var p models.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdateProductWithManagerCheck(id, p, manager.ID); err != nil {
		if strings.Contains(err.Error(), "unique_product_name") {
			http.Error(w, "product name must be unique", http.StatusConflict)
			return
		}
		http.Error(w, "product not found or access denied", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	login := r.Context().Value("manager_login").(string)
	manager, err := h.repo.GetManagerByLogin(login)
	if err != nil {
		http.Error(w, "manager not found", http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteProductWithManagerCheck(id, manager.ID); err != nil {
		http.Error(w, "product not found or access denied", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
