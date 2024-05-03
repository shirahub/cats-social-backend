package service

import (
	"app/domain"
	"app/port"
	"context"
	"time"
)

type catMatchSvc struct {
	cRepo port.CatRepository
	mRepo port.CatMatchRepository
}

func NewCatMatchService(catRepo port.CatRepository, matchRepo port.CatMatchRepository) *catMatchSvc {
	return &catMatchSvc{catRepo, matchRepo}
}

func (s *catMatchSvc) Create(c context.Context, catMatch *domain.CatMatch) (*domain.CatMatch, error) {
	issuerCat, err := s.cRepo.GetById(c, catMatch.IssuerCatId)
	if err != nil {
		return nil, err
	}
	if issuerCat.HasMatched {
		return nil, domain.ErrMatchWithTaken
	}
	receiverCat, err := s.cRepo.GetById(c, catMatch.ReceiverCatId)
	if err != nil {
		return nil, err
	}
	if receiverCat.HasMatched {
		return nil, domain.ErrMatchWithTaken
	}
	if issuerCat.UserId == receiverCat.UserId {
		return nil, domain.ErrMatchWithOwnedCat
	}
	if issuerCat.Sex == receiverCat.Sex {
		return nil, domain.ErrMatchWithSameSex
	}

	return s.mRepo.Create(c, catMatch)
}

func (s *catMatchSvc) List(c context.Context, userId string) ([]domain.CatMatchDetail, error) {
	return s.mRepo.List(c, userId)
}

func (s *catMatchSvc) Approve(c context.Context, matchId string, userId string) (id string, updatedAt time.Time, err error) {
	id, updatedAt, err = s.mRepo.ApproveAndInvalidateOthers(c, matchId, userId)
	if err == domain.ErrNotFound {
		_, err = s.mRepo.GetReceivedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return id, updatedAt, err
}

func (s *catMatchSvc) Reject(c context.Context, matchId string, userId string) (id string, updatedAt time.Time, err error) {
	id, updatedAt, err = s.mRepo.Reject(c, matchId, userId)
	if err == domain.ErrNotFound {
		_, err = s.mRepo.GetReceivedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return id, updatedAt, err
}

func (s *catMatchSvc) Delete(c context.Context, matchId string, userId string) (string, time.Time, error) {
	matchId, deletedAt, err := s.mRepo.Delete(c, matchId, userId)
	if err == domain.ErrNotFound {
		_, err = s.mRepo.GetIssuedByIdUserId(matchId, userId)
		if err == domain.ErrNotFound {
			return "", time.Time{}, domain.ErrNotFound
		}
		return "", time.Time{}, domain.ErrMatchResponded
	}
	return matchId, deletedAt, nil
}
