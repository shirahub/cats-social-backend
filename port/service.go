package port

import (
	"app/domain"
	"context"
	"time"
)

type CatManagementService interface {
	Create(context.Context, *domain.CreateCatRequest) (*domain.Cat, error)
	List(context.Context, *domain.GetCatsRequest) ([]domain.Cat, error)
	Update(*domain.Cat) (*domain.Cat, error)
	Delete(catId string, userId string, ) (string, time.Time, error)
}

type CatMatchService interface {
	Create(catMatch *domain.CatMatch) (*domain.CatMatch, error)
	List() ([]domain.CatMatch, error)
	Approve(matchId string, userId string) (id string, updatedAt time.Time, err error)
	Reject(matchId string, userId string) (id string, updatedAt time.Time, err error)
	Delete(matchId string, userId string) (id string, deletedAt time.Time, err error)
}