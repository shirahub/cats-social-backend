package port

import "app/domain"

type CatManagementService interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	List(*domain.GetCatsRequest) ([]domain.Cat, error)
	Update(*domain.Cat) (*domain.Cat, error)
	Delete(userId string, catId string) (string, string, error)
}

type CatMatchService interface {
	Create(catMatch *domain.CatMatch) (*domain.CatMatch, error)
	Delete(userId string, matchId string) (id string, deletedAt string, err error)
}