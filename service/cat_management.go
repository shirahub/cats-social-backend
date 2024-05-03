package service

import (
	"app/domain"
	"app/port"
	"context"
	"time"
)

type catManagementSvc struct {
	cRepo port.CatRepository
	mRepo port.CatMatchRepository
}

func NewCatManagementService(cRepo port.CatRepository, mRepo port.CatMatchRepository) *catManagementSvc {
	return &catManagementSvc{cRepo, mRepo}
}

func (h *catManagementSvc) Create(ctx context.Context, cat *domain.CreateCatRequest) (*domain.Cat, error) {
	return h.cRepo.Create(ctx, cat)
}

func (h *catManagementSvc) List(ctx context.Context, req *domain.GetCatsRequest) ([]domain.Cat, error) {
	return h.cRepo.List(ctx, req)
}

func (h *catManagementSvc) Update(ctx context.Context, updatedCat *domain.Cat) (*domain.Cat, error) {
	cat, err := h.cRepo.GetById(ctx, updatedCat.Id)
	if err != nil {
		return nil, err
	}
	if cat.Sex != updatedCat.Sex {
		IsCatInAnyMatch, err := h.mRepo.IsCatInAnyMatch(ctx, updatedCat.Id)
		if err != nil {
			return nil, err
		}
		if IsCatInAnyMatch {
			return nil, domain.ErrCatInMatch
		}
	}
	return h.cRepo.Update(ctx, updatedCat)
}

func (h *catManagementSvc) Delete(ctx context.Context, catId string, userId string) (string, time.Time, error) {
	return h.cRepo.Delete(ctx, catId, userId)
}