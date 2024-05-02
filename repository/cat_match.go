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

const getIssuedByIdUserIdQuery = `
	SELECT cat_matches.id, message, issuer_cat_id, receiver_cat_id, status, cat_matches.created_at
	FROM cat_matches
	INNER JOIN cats ON cat_matches.issuer_cat_id = cats.id
	WHERE cat_matches.id = $1 AND user_id = $2 AND cat_matches.deleted_at is null
`

const getReceivedByIdUserIdQuery = `
	SELECT cat_matches.id, message, issuer_cat_id, receiver_cat_id, status, cat_matches.created_at
	FROM cat_matches
	INNER JOIN cats ON cat_matches.receiver_cat_id = cats.id
	WHERE cat_matches.id = $1 AND user_id = $2 AND cat_matches.deleted_at is null
`

const updateStatusMatchQuery = `
	UPDATE cat_matches
	SET status = $1
	FROM cats
	WHERE user_id = $2
	AND receiver_cat_id = cats.id
	AND cat_matches.id = $3
	AND status = 'pending'
	AND cat_matches.deleted_at is null
	RETURNING cat_matches.id, cat_matches.updated_at
`

const invalidateMatchesQuery = `
	UPDATE cat_matches
	SET status = 'invalid'
	FROM cats
	WHERE status = 'pending'
	AND cat_matches.deleted_at is null
	AND (
		issuer_cat_id = $1 OR receiver_cat_id = $1
		OR issuer_cat_id = $2 OR receiver_cat_id = $2
	)
`

const updateDeletedAtMatchQuery = `
	UPDATE cat_matches
	SET deleted_at = NOW()
	FROM cats
	WHERE user_id = $1
	AND issuer_cat_id = cats.id
	AND cat_matches.id = $2
	AND status = 'pending'
	AND cat_matches.deleted_at is null
	RETURNING cat_matches.id, cat_matches.deleted_at
`

func (r *CatMatchRepo) Create(match *domain.CatMatch) (*domain.CatMatch, error) {
	err := r.db.QueryRow(
		createMatchQuery, match.Message, match.IssuerCatId, match.ReceiverCatId, "pending",
	).Scan(&match.Id, &match.CreatedAt)
	return match, err
}

func (r *CatMatchRepo) List() ([]domain.CatMatch, error) {
	return nil, nil
}

func (r *CatMatchRepo) GetIssuedByIdUserId(matchId string, userId string) (*domain.CatMatch, error) {
	var match domain.CatMatch
	err := r.db.QueryRow(
		getIssuedByIdUserIdQuery,
		matchId, userId,
	).Scan(&match.Id, &match.Message, &match.IssuerCatId, &match.ReceiverCatId, &match.Status, &match.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return &match, err
}

func (r *CatMatchRepo) GetReceivedByIdUserId(matchId string, userId string) (*domain.CatMatch, error) {
	var match domain.CatMatch
	err := r.db.QueryRow(
		getReceivedByIdUserIdQuery,
		matchId, userId,
	).Scan(&match.Id, &match.Message, &match.IssuerCatId, &match.ReceiverCatId, &match.Status, &match.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrNotFound
	}
	return &match, err
}

func (r *CatMatchRepo) ApproveAndInvalidateOthers(matchId string, userId string) (string, time.Time, error) {
	var updatedMatchId string
	var updatedAt time.Time
	err := r.db.QueryRow(updateStatusMatchQuery, "approved", userId, matchId).Scan(&updatedMatchId, &updatedAt)
	if err == sql.ErrNoRows {
		return "", time.Time{}, domain.ErrNotFound
	}
	return updatedMatchId, updatedAt, err
}

func (r *CatMatchRepo) Reject(matchId string, userId string) (string, time.Time, error) {
	var updatedMatchId string
	var updatedAt time.Time
	err := r.db.QueryRow(updateStatusMatchQuery, "rejected", userId, matchId).Scan(&updatedMatchId, &updatedAt)
	if err == sql.ErrNoRows {
		return "", time.Time{}, domain.ErrNotFound
	}
	return updatedMatchId, updatedAt, err
}

func (r *CatMatchRepo) Delete(matchId string, userId string) (string, time.Time, error) {
	var deletedMatchId string
	var deletedAt time.Time
	err := r.db.QueryRow(updateDeletedAtMatchQuery, userId, matchId).Scan(&deletedMatchId, &deletedAt)
	if err == sql.ErrNoRows {
		return "", time.Time{}, domain.ErrNotFound
	}
	return deletedMatchId, deletedAt, err
}
