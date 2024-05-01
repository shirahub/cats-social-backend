package port

import (
	"app/domain"
	"time"
)

type CatRepository interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	List(*domain.GetCatsRequest) ([]domain.Cat, error)
	Update(*domain.Cat) (*domain.Cat, error)
	Delete(catId string, userId string) (id string, deletedAt time.Time, err error)
}

type CatMatchRepository interface {
	Create(*domain.CatMatch) (*domain.CatMatch, error)
	GetByIdUserId(matchId string, userId string) (*domain.CatMatch, error)
	Delete(matchId string, userId string) (id string, deletedAt time.Time, err error)
}