package port

import "app/domain"

type CatRepository interface {
	Create(*domain.CreateCatRequest) error
	Update(*domain.Cat) error
}

type MatchRepository interface {
}