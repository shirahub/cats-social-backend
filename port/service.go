package port

import "app/domain"

type CatManagementService interface {
	Create(*domain.CreateCatRequest) error
}