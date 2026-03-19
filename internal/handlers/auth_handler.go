package handlers

import (
	"api-pedidos/internal/services"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.error(w, http.StatusBadRequest, "payload inválido")
		return
	}
	if req.Email == "" || req.Password == "" {
		h.error(w, http.StatusBadRequest, "email e senha são obrigatórios")
		return
	}
	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		h.error(w, http.StatusUnauthorized, err.Error())
		return
	}
	h.json(w, http.StatusOK, map[string]string{"token": token})
}

func (h *AuthHandler) json(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func (h *AuthHandler) error(w http.ResponseWriter, status int, msg string) {
	h.json(w, status, map[string]string{"error": msg})
}
