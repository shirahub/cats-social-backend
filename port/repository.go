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
	GetIssuedByIdUserId(matchId string, userId string) (*domain.CatMatch, error)
	GetReceivedByIdUserId(matchId string, userId string) (*domain.CatMatch, error)
	UpdateStatus(matchId string, userId string, status string) (id string, updatedAt time.Time, err error)
	Invalidate(cat1Id string, cat2Id string) error
	Delete(matchId string, userId string) (id string, deletedAt time.Time, err error)
}