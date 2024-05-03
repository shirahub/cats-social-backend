package repository

import (
	"app/domain"
	"database/sql"
	"github.com/lib/pq"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(user domain.User) (string, error) {
	var userId string
	err := r.db.QueryRow(
		"INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id",
		user.Email, user.Name, user.Password,
	).Scan(&userId)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" {
			return "", domain.ErrEmailTaken
		}
		return "", err
	}
	return userId, nil
}

func (r *UserRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return &user, err
}