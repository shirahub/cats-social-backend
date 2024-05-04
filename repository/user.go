package repository

import (
	"app/domain"
	"context"
	"database/sql"
	"github.com/lib/pq"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(c context.Context, user domain.User) (string, error) {
	var userId string
	err := r.db.QueryRowContext(
		c, "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id",
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

func (r *UserRepo) FindByEmail(c context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.db.QueryRowContext(c, "SELECT * FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return &user, err
}