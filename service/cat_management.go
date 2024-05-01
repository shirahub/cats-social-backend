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

func (h *catManagementSvc) Create(cat *domain.CreateCatRequest) error {
	return h.repo.Create(cat)
}

func (h *catManagementSvc) Update(cat *domain.Cat) error {
	return h.repo.Update(cat)
}