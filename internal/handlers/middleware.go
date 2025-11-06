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

		// Сохраняем login в контексте
		ctx := context.WithValue(r.Context(), "manager_login", login)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
