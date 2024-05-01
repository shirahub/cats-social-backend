package repository

import (
	"app/domain"
	"fmt"
	"strings"
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

const getQuery = `
	SELECT id, name, race, sex, age_in_month, description, image_urls, user_id, created_at
	FROM cats
	WHERE deleted_at is null
	?
	ORDER BY created_at DESC
`

const updateQuery = `
	UPDATE cats
	SET name = $1, race = $2, sex = $3, age_in_month = $4, description = $5, image_urls = $6
	WHERE user_id = $7 and id = $8 and deleted_at is null
	RETURNING id, updated_at
`

const updateDeletedAtQuery = `
	UPDATE cats
	SET deleted_at = NOW()
	WHERE user_id = $1 and id = $2 and deleted_at is null
	RETURNING id, deleted_at
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

func formGetQuery(req *domain.GetCatsRequest) (string, []interface{}) {
	var query strings.Builder
	filterArgs := make([]interface{}, 0)
	if req.Id != "" {
		filterArgs = append(filterArgs, req.Id)
		query.WriteString(fmt.Sprintf(" AND id = $%d", len(filterArgs)))
	}
	if req.Race != "" {
		filterArgs = append(filterArgs, req.Race)
		query.WriteString(fmt.Sprintf(" AND race = $%d", len(filterArgs)))
	}
	if req.Sex != "" {
		filterArgs = append(filterArgs, req.Sex)
		query.WriteString(fmt.Sprintf(" AND sex = $%d", len(filterArgs)))
	}
	if req.UserId != "" {
		filterArgs = append(filterArgs, req.UserId)
		query.WriteString(fmt.Sprintf(" AND user_id = $%d", len(filterArgs)))
	}
	return strings.Replace(getQuery, "?", query.String(), 1), filterArgs
}

func (r *CatRepo) List(req *domain.GetCatsRequest) ([]domain.Cat, error) {
	query, filterArgs := formGetQuery(req)
	rows, err := r.db.Query(query, filterArgs...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cats []domain.Cat
	for rows.Next() {
		var cat domain.Cat
		if err := rows.Scan(
			&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description,
			pq.Array(&cat.ImageUrls), &cat.UserId, &cat.CreatedAt,
		); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cats, err
}

func (r *CatRepo) Update(cat *domain.Cat) (*domain.Cat, error) {
	err := r.db.QueryRow(
		updateQuery,
		cat.Name, cat.Race, cat.Sex, cat.AgeInMonth, cat.Description, pq.Array(cat.ImageUrls),
		cat.UserId, cat.Id,
	).Scan(&cat.Id, &cat.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		} else {
			return nil, err
		}
	}
	return cat, err
}

func (r *CatRepo) Delete(userId string, catId string) (string, string, error) {
	var deletedCatId, deletedAt string
	err := r.db.QueryRow(
		updateDeletedAtQuery,
		userId, catId,
	).Scan(&deletedCatId, &deletedAt)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return "", "", domain.ErrNotFound
		} else {
			return "", "", err
		}
	}
	
	return deletedCatId, deletedAt, err
}