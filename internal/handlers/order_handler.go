package handlers

import (
	"api-pedidos/internal/models"
	"api-pedidos/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req models.Order
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.error(w, http.StatusBadRequest, "payload inválido")
		return
	}
	if err := h.service.Create(&req); err != nil {
		h.error(w, http.StatusBadRequest, err.Error())
		return
	}
	h.json(w, http.StatusCreated, req)
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.service.List()
	if err != nil {
		h.error(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.json(w, http.StatusOK, orders)
}

func (h *OrderHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	orderIDStr := chi.URLParam(r, "orderID")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		h.error(w, http.StatusBadRequest, "orderID inválido")
		return
	}
	var payload struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		h.error(w, http.StatusBadRequest, "payload inválido")
		return
	}
	if payload.Status == "" {
		h.error(w, http.StatusBadRequest, "status é obrigatório")
		return
	}
	if err := h.service.UpdateStatus(orderID, payload.Status); err != nil {
		h.error(w, http.StatusBadRequest, err.Error())
		return
	}
	h.json(w, http.StatusOK, map[string]string{"message": "status atualizado"})
}

func (h *OrderHandler) json(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func (h *OrderHandler) error(w http.ResponseWriter, status int, msg string) {
	h.json(w, status, map[string]string{"error": msg})
}
