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

func (r *UserRepo) Create(user domain.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (email, name, password) VALUES ($1, $2, $3)",
		user.Email, user.Name, user.Password,
	)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" {
			return domain.ErrEmailTaken
		}
		return err
	}
	return err
}

func (r *UserRepo) FindByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email)
	user := domain.User{}
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}