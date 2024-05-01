package port

import "app/domain"

type CatRepository interface {
	Create(*domain.CreateCatRequest) error
}

type MatchRepository interface {
}