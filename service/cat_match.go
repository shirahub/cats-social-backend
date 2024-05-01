package service

import (
	"app/domain"
	"app/port"
	"time"
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

func (s *catMatchSvc) Delete(userId string, catMatchId string) (string, time.Time, error) {
	return s.repo.Delete(userId, catMatchId)
}