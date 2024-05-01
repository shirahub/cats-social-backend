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

func (s *catMatchSvc) Approve(matchId string, userId string) (id string, updatedAt time.Time, err error) {
	id, updatedAt, err = s.repo.UpdateStatus(matchId, userId, "approved")
	if err == domain.ErrNotFound {
		_, err = s.repo.GetReceivedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return id, updatedAt, err
}

func (s *catMatchSvc) Reject(matchId string, userId string) (id string, updatedAt time.Time, err error) {
	id, updatedAt, err = s.repo.UpdateStatus(matchId, userId, "rejected")
	if err == domain.ErrNotFound {
		_, err = s.repo.GetReceivedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return id, updatedAt, err
}

func (s *catMatchSvc) Delete(catMatchId string, userId string) (string, time.Time, error) {
	matchId, deletedAt, err := s.repo.Delete(userId, catMatchId)
	if err == domain.ErrNotFound {
		_, err = s.repo.GetIssuedByIdUserId(catMatchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return matchId, deletedAt, nil
}