package port

import "app/domain"

type CatRepository interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	List(*domain.GetCatsRequest) ([]domain.Cat, error)
	Update(*domain.Cat) (*domain.Cat, error)
	Delete(userId string, catId string) (id string, deletedAt string, err error)
}

type CatMatchRepository interface {
	Create(*domain.CatMatch) (*domain.CatMatch, error)
}