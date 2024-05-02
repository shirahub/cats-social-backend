package service

import (
	"app/domain"
	"app/port"
	"context"
	"time"
)

type catManagementSvc struct {
	repo port.CatRepository
}

func NewCatManagementService(repo port.CatRepository) *catManagementSvc {
	return &catManagementSvc{repo}
}

func (h *catManagementSvc) Create(ctx context.Context, cat *domain.CreateCatRequest) (*domain.Cat, error) {
	return h.repo.Create(ctx, cat)
}

func (h *catManagementSvc) List(ctx context.Context, req *domain.GetCatsRequest) ([]domain.Cat, error) {
	return h.repo.List(ctx, req)
}

func (h *catManagementSvc) Update(cat *domain.Cat) (*domain.Cat, error) {
	return h.repo.Update(cat)
}

func (h *catManagementSvc) Delete(catId string, userId string) (string, time.Time, error) {
	return h.repo.Delete(catId, userId)
}