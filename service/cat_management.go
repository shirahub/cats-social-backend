package service

import (
	"app/domain"
	"app/port"
)


type catManagementSvc struct {
	repo port.CatRepository
}

func NewCatManagementService(repo port.CatRepository) *catManagementSvc {
	return &catManagementSvc{repo}
}

func (h *catManagementSvc) Create(cat *domain.CreateCatRequest) (*domain.Cat, error) {
	return h.repo.Create(cat)
}

func (h *catManagementSvc) Update(cat *domain.Cat) error {
	return h.repo.Update(cat)
}

func (h *catManagementSvc) Delete(userId string, catId string) error {
	return h.repo.Delete(userId, catId)
}