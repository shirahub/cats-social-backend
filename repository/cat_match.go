package repository

import (
	"app/domain"
	"database/sql"
)

type CatMatchRepo struct {
	db *sql.DB
}

func NewCatMatchRepo(db *sql.DB) *CatMatchRepo {
	return &CatMatchRepo{db}
}

const createMatchQuery = `
	INSERT INTO cat_matches
	(message, issuer_cat_id, receiver_cat_id, status)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
`

func (r *CatMatchRepo) Create(match *domain.CatMatch) (*domain.CatMatch, error) {
	err := r.db.QueryRow(
		createMatchQuery, match.Message, match.IssuerCatId, match.ReceiverCatId, "pending",
	).Scan(&match.Id, &match.CreatedAt)
	return match, err
}
