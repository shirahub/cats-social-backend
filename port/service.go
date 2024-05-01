package port

import "app/domain"

type CatManagementService interface {
	Create(*domain.CreateCatRequest) (*domain.Cat, error)
	Update(*domain.Cat) (*domain.Cat, error)
	Delete(userId string, catId string) (string, string, error)
}