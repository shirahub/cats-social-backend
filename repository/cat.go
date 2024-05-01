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

const createQuery = `
	INSERT INTO cats 
	(name, race, sex, age_in_month, description, image_urls, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at
`

func (r *CatRepo) Create(cat *domain.CreateCatRequest) (*domain.Cat, error) {
	newRecord := domain.Cat{
		Name: cat.Name,
		Race: cat.Race,
		Sex: cat.Sex,
		AgeInMonth: cat.AgeInMonth,
		Description: cat.Description,
		ImageUrls: cat.ImageUrls,
		UserId: cat.UserId,
	}
	err := r.db.QueryRow(
		createQuery, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth,
		cat.Description, pq.Array(cat.ImageUrls), cat.UserId,
	).Scan(&newRecord.Id, &newRecord.CreatedAt)
	fmt.Println(err)
	return &newRecord, err
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

func (r *CatRepo) Delete(userId string, catId string) error {
	_, err := r.db.Exec(
		`DELETE FROM cats
		WHERE user_id = $1 and id = $2`,
		userId, catId,
	)
	fmt.Println(err)
	return err
}