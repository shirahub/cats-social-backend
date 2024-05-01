package port

import "app/domain"

type CatManagementService interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	Update(*domain.Cat) error
}