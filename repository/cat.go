package repository

import (
	"database/sql"
)

type CatRepo struct {
	db *sql.DB
}

func NewCatRepo(db *sql.DB) *CatRepo {
	return &CatRepo{db}
}