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
		"INSERT INTO cats (name, race, sex, age_in_month, description, image_urls, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description,
		pq.Array(cat.ImageUrls), cat.UserId,
	)
	fmt.Println(err)
	return err
}

func (r *CatRepo) Update(cat *domain.Cat) error {
	_, err := r.db.Exec(
		`UPDATE cats
		SET name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, image_urls = $6
		WHERE user_id = $7 and id = $8`,
		cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, pq.Array(cat.ImageUrls),
		cat.UserId, cat.Id,
	)
	fmt.Println(err)
	return err
}