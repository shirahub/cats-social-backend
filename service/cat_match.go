package service

import (
	"app/domain"
	"app/port"
	"context"
	"time"
)

type catMatchSvc struct {
	repo port.CatMatchRepository
}

func NewCatMatchService(repo port.CatMatchRepository) *catMatchSvc {
	return &catMatchSvc{repo}
}

func (s *catMatchSvc) Create(c context.Context, catMatch *domain.CatMatch) (*domain.CatMatch, error) {
	return s.repo.Create(c, catMatch)
}

func (s *catMatchSvc) List() ([]domain.CatMatch, error) {
	return nil, nil
}

func (s *catMatchSvc) Approve(matchId string, userId string) (id string, updatedAt time.Time, err error) {
	id, updatedAt, err = s.repo.ApproveAndInvalidateOthers(matchId, userId)
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
	id, updatedAt, err = s.repo.Reject(matchId, userId)
	if err == domain.ErrNotFound {
		_, err = s.repo.GetReceivedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return id, updatedAt, err
}

func (s *catMatchSvc) Delete(c context.Context, catMatchId string, userId string) (string, time.Time, error) {
	matchId, deletedAt, err := s.repo.Delete(c, userId, catMatchId)
	if err == domain.ErrNotFound {
		_, err = s.repo.GetIssuedByIdUserId(catMatchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return matchId, deletedAt, nil
}
