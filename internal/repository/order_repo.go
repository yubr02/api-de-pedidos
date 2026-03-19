package repository

import (
	"api-pedidos/internal/models"
	"database/sql"
	"fmt"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	query := `INSERT INTO orders (customer_name, items, total, status) VALUES (?, ?, ?, ?)`
	res, err := r.db.Exec(query, order.CustomerName, order.Items, order.Total, order.Status)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	order.ID = int(id)
	return nil
}

func (r *OrderRepository) List() ([]models.Order, error) {
	rows, err := r.db.Query(`SELECT id, customer_name, items, total, status, created_at FROM orders ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.Items, &o.Total, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) UpdateStatus(orderID int, status string) error {
	res, err := r.db.Exec(`UPDATE orders SET status = ? WHERE id = ?`, status, orderID)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("order not found")
	}
	return nil
}
