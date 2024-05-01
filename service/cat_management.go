package service

import (
	"app/domain"
	"app/port"
	"time"
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

func (h *catManagementSvc) List(req *domain.GetCatsRequest) ([]domain.Cat, error) {
	return h.repo.List(req)
}

func (h *catManagementSvc) Update(cat *domain.Cat) (*domain.Cat, error) {
	return h.repo.Update(cat)
}

func (h *catManagementSvc) Delete(catId string, userId string) (string, time.Time, error) {
	return h.repo.Delete(catId, userId)
}