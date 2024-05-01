package repository

import (
	"app/domain"
	"fmt"
	"database/sql"
	"github.com/lib/pq"
)

type CatRepo struct {
	db *sql.DB
}

func NewCatRepo(db *sql.DB) *CatRepo {
	return &CatRepo{db}
}

func (r *CatRepo) Create(cat *domain.CreateCatRequest) error {
	_, err := r.db.Exec(
		"INSERT INTO cats (name, race, sex, age_in_month, description, image_urls) VALUES ($1, $2, $3, $4, $5, $6)",
		cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, pq.Array(cat.ImageUrls),
	)
	fmt.Println(err)
	return err
}