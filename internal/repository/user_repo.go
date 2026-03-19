package repository

import (
	"api-pedidos/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var u models.User
	query := `SELECT id, email, password, fullname FROM users WHERE email = ? LIMIT 1`
	row := r.db.QueryRow(query, email)
	if err := row.Scan(&u.ID, &u.Email, &u.Password, &u.FullName); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
