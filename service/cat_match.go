package service

import (
	"app/domain"
	"app/port"
)


type catMatchSvc struct {
	repo port.CatMatchRepository
}

func NewCatMatchService(repo port.CatMatchRepository) *catMatchSvc {
	return &catMatchSvc{repo}
}

func (s *catMatchSvc) Create(catMatch *domain.CatMatch) (*domain.CatMatch, error) {
	return s.repo.Create(catMatch)
}
