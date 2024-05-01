package port

import (
	"app/domain"
	"time"
)

type CatManagementService interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	List(*domain.GetCatsRequest) ([]domain.Cat, error)
	Update(*domain.Cat) (*domain.Cat, error)
	Delete(userId string, catId string) (string, time.Time, error)
}

type CatMatchService interface {
	Create(catMatch *domain.CatMatch) (*domain.CatMatch, error)
	Delete(userId string, matchId string) (id string, deletedAt time.Time, err error)
}