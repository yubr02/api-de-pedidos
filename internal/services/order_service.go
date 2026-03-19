package services

import (
	"api-pedidos/internal/models"
	"api-pedidos/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(order *models.Order) error {
	if order.CustomerName == "" {
		return &ServiceError{Message: "customer_name is required"}
	}
	if order.Items == "" {
		return &ServiceError{Message: "items are required"}
	}
	if order.Total <= 0 {
		return &ServiceError{Message: "total must be greater than zero"}
	}
	if order.Status == "" {
		order.Status = "pending"
	}
	return s.repo.Create(order)
}

func (s *OrderService) List() ([]models.Order, error) {
	return s.repo.List()
}

func (s *OrderService) UpdateStatus(orderID int, status string) error {
	if orderID <= 0 {
		return &ServiceError{Message: "order ID inválido"}
	}
	if status == "" {
		return &ServiceError{Message: "status é obrigatório"}
	}
	return s.repo.UpdateStatus(orderID, status)
}
