package repository

import (
	"time"
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

const deleteMatchQuery = `
	DELETE FROM cat_matches
	WHERE id = $1 AND status = 'pending'
	RETURNING id
`

func (r *CatMatchRepo) Create(match *domain.CatMatch) (*domain.CatMatch, error) {
	err := r.db.QueryRow(
		createMatchQuery, match.Message, match.IssuerCatId, match.ReceiverCatId, "pending",
	).Scan(&match.Id, &match.CreatedAt)
	return match, err
}

func (r *CatMatchRepo) Delete(userId string, matchId string) (string, string, error) {
	var deletedMatchId string
	err := r.db.QueryRow(deleteMatchQuery, matchId).Scan(&deletedMatchId)
	return deletedMatchId, time.Now().String(), err
}
