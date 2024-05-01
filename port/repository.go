package port

import "app/domain"

type CatRepository interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	Update(*domain.Cat) error
	Delete(userId string, catId string) error
}

type MatchRepository interface {
}