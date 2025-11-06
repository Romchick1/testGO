package handlers

import (
	"context"
	"net/http"
)

func (h *ManagerHandler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := r.Header.Get("X-Manager-Login")
		if login == "" {
			http.Error(w, "X-Manager-Login header is required", http.StatusBadRequest)
			return
		}

		_, err := h.repo.GetManagerByLogin(login)
		if err != nil {
			http.Error(w, "manager not found", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "manager_login", login)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
