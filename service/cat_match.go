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

func (s *catMatchSvc) List(c context.Context) ([]domain.CatMatchDetail, error) {
	return s.repo.List(c)
}

func (s *catMatchSvc) Approve(c context.Context, matchId string, userId string) (id string, updatedAt time.Time, err error) {
	id, updatedAt, err = s.repo.ApproveAndInvalidateOthers(c, matchId, userId)
	if err == domain.ErrNotFound {
		_, err = s.repo.GetReceivedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return id, updatedAt, err
}

func (s *catMatchSvc) Reject(c context.Context, matchId string, userId string) (id string, updatedAt time.Time, err error) {
	id, updatedAt, err = s.repo.Reject(c, matchId, userId)
	if err == domain.ErrNotFound {
		_, err = s.repo.GetReceivedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return id, updatedAt, err
}

func (s *catMatchSvc) Delete(c context.Context, matchId string, userId string) (string, time.Time, error) {
	matchId, deletedAt, err := s.repo.Delete(c, matchId, userId)
	if err == domain.ErrNotFound {
		_, err = s.repo.GetIssuedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return matchId, deletedAt, nil
}
