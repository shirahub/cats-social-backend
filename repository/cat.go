package repository

import (
	"app/domain"
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"
	"strconv"
	"database/sql"
	"github.com/lib/pq"
)

type CatRepo struct {
	db *sql.DB
}

func NewCatRepo(db *sql.DB) *CatRepo {
	return &CatRepo{db}
}

const createCatQuery = `
	INSERT INTO cats 
	(name, race, sex, age_in_month, description, image_urls, user_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at
`

const getByIdQuery = `
	SELECT id, sex, user_id, has_matched
	FROM cats
	WHERE id = $1 and deleted_at is null
`

const getQuery = `
	SELECT id, name, race, sex, age_in_month, description, image_urls, user_id, has_matched, created_at
	FROM cats
	WHERE deleted_at is null
	{{ where_query }}
	ORDER BY created_at DESC
	{{ pagination_query }}
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

const comparisonPattern = `^([<>]=?|=)\s*(\d+)$`

func isValidId(id string) bool {
	_, err := strconv.ParseInt(id, 10, 64)
	return err == nil
}

func (r *CatRepo) Create(ctx context.Context, cat *domain.CreateCatRequest) (*domain.Cat, error) {
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
		createCatQuery, cat.Name, cat.Race, cat.Sex, cat.AgeInMonth,
		cat.Description, pq.Array(cat.ImageUrls), cat.UserId,
	).Scan(&newRecord.Id, &newRecord.CreatedAt)
	return &newRecord, err
}

func (r *CatRepo) GetById(c context.Context, id string) (*domain.Cat, error) {
	if !isValidId(id) {
		return nil, domain.ErrNotFound
	}

	var cat domain.Cat
	err := r.db.QueryRowContext(c, getByIdQuery, id).
		Scan(&cat.Id, &cat.Sex, &cat.UserId, &cat.HasMatched)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &cat, nil
}

func formGetQuery(req *domain.GetCatsRequest) (string, []interface{}) {
	var whereQuery, paginationQuery strings.Builder
	filterArgs := make([]interface{}, 0)
	if req.Id != "" {
		filterArgs = append(filterArgs, req.Id)
		whereQuery.WriteString(fmt.Sprintf(" AND id = $%d", len(filterArgs)))
	}
	if req.Race != "" {
		filterArgs = append(filterArgs, req.Race)
		whereQuery.WriteString(fmt.Sprintf(" AND race = $%d", len(filterArgs)))
	}
	if req.Sex != "" {
		filterArgs = append(filterArgs, req.Sex)
		whereQuery.WriteString(fmt.Sprintf(" AND sex = $%d", len(filterArgs)))
	}
	if req.HasMatched != nil {
		filterArgs = append(filterArgs, req.HasMatched)
		whereQuery.WriteString(fmt.Sprintf(" AND has_matched = $%d", len(filterArgs)))
	}
	if req.AgeInMonth != "" {
		re := regexp.MustCompile(comparisonPattern)
		matches := re.FindStringSubmatch(req.AgeInMonth)
		filterArgs = append(filterArgs, matches[2])
		whereQuery.WriteString(fmt.Sprintf(" AND age_in_month %s $%d", matches[1], len(filterArgs)))
	}
	if req.UserId != "" && req.Owned != nil {
		filterArgs = append(filterArgs, req.UserId)
		if *req.Owned {
			whereQuery.WriteString(fmt.Sprintf(" AND user_id = $%d", len(filterArgs)))
		} else {
			whereQuery.WriteString(fmt.Sprintf(" AND user_id != $%d", len(filterArgs)))
		}
	}
	if req.Name != "" {
		filterArgs = append(filterArgs, req.Name)
		whereQuery.WriteString(fmt.Sprintf(" AND name = $%d", len(filterArgs)))
	}

	fullQuery := strings.Replace(getQuery, "{{ where_query }}", whereQuery.String(), 1)

	filterArgs = append(filterArgs, req.Limit)
	paginationQuery.WriteString(fmt.Sprintf("LIMIT $%d", len(filterArgs)))
	filterArgs = append(filterArgs, req.Offset)
	paginationQuery.WriteString(fmt.Sprintf(" OFFSET $%d", len(filterArgs)))

	fullQuery = strings.Replace(fullQuery, "{{ pagination_query }}", paginationQuery.String(), 1)
	return fullQuery, filterArgs
}

func (r *CatRepo) List(c context.Context, req *domain.GetCatsRequest) ([]domain.Cat, error) {
	var cats = make([]domain.Cat, 0)
	if req.UserId != "" {
		if !isValidId(req.UserId) {
			return nil, domain.ErrNotFound
		}
	}
	
	query, filterArgs := formGetQuery(req)
	rows, err := r.db.QueryContext(c, query, filterArgs...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var cat domain.Cat
		if err := rows.Scan(
			&cat.Id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description,
			pq.Array(&cat.ImageUrls), &cat.UserId, &cat.HasMatched, &cat.CreatedAt,
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

func (r *CatRepo) Update(c context.Context, cat *domain.Cat) (*domain.Cat, error) {
	if !isValidId(cat.Id) || !isValidId(cat.UserId) {
		return nil, domain.ErrNotFound
	}
	err := r.db.QueryRowContext(
		c, updateQuery,
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

func (r *CatRepo) Delete(c context.Context, catId string, userId string) (string, time.Time, error) {
	if !isValidId(catId) || !isValidId(userId) {
		return "", time.Time{}, domain.ErrNotFound
	}

	var deletedCatId string
	var deletedAt time.Time
	err := r.db.QueryRowContext(
		c, updateDeletedAtQuery, userId, catId,
	).Scan(&deletedCatId, &deletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", time.Time{}, domain.ErrNotFound
		} else {
			return "", time.Time{}, err
		}
	}
	
	return deletedCatId, deletedAt, err
}