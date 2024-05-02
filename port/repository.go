package port

import (
	"app/domain"
	"context"
	"time"
)

type CatRepository interface {
	Create(context.Context, *domain.CreateCatRequest) (*domain.Cat, error)
	List(context.Context, *domain.GetCatsRequest) ([]domain.Cat, error)
	Update(context.Context, *domain.Cat) (*domain.Cat, error)
	Delete(c context.Context, catId string, userId string) (id string, deletedAt time.Time, err error)
}

type CatMatchRepository interface {
	Create(context.Context, *domain.CatMatch) (*domain.CatMatch, error)
	List() ([]domain.CatMatch, error)
	GetIssuedByIdUserId(matchId string, userId string) (*domain.CatMatch, error)
	GetReceivedByIdUserId(matchId string, userId string) (*domain.CatMatch, error)
	ApproveAndInvalidateOthers(matchId string, receiverUserId string) (id string, updatedAt time.Time, err error)
	Reject(matchId string, receiverUserId string) (id string, updatedAt time.Time, err error)
	Delete(c context.Context, matchId string, userId string) (id string, deletedAt time.Time, err error)
}