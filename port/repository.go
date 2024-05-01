package port

import "app/domain"

type CatRepository interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	Update(*domain.Cat) error
}

type MatchRepository interface {
}