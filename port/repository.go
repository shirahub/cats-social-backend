package port

import (
	"app/domain"
	"context"
	"time"
)

type CatRepository interface {
	Create(context.Context, *domain.CreateCatRequest) (*domain.Cat, error)
	GetById(context.Context, string) (*domain.Cat, error)
	List(context.Context, *domain.GetCatsRequest) ([]domain.Cat, error)
	Update(context.Context, *domain.Cat) (*domain.Cat, error)
	Delete(c context.Context, catId string, userId string) (id string, deletedAt time.Time, err error)
}

type CatMatchRepository interface {
	Create(context.Context, *domain.CatMatch) (*domain.CatMatch, error)
	List(c context.Context, userId string) ([]domain.CatMatchDetail, error)
	GetIssuedByIdUserId(c context.Context, matchId string, userId string) (*domain.CatMatch, error)
	GetReceivedByIdUserId(c context.Context, matchId string, userId string) (*domain.CatMatch, error)
	IsCatInAnyMatch(c context.Context, catId string) (bool, error)
	AnyMatchExists(c context.Context, cat1Id, cat2Id string) (bool, error)
	ApproveAndInvalidateOthers(c context.Context, matchId string, receiverUserId string) (id string, updatedAt time.Time, err error)
	Reject(c context.Context, matchId string, receiverUserId string) (id string, updatedAt time.Time, err error)
	Delete(c context.Context, matchId string, userId string) (id string, deletedAt time.Time, err error)
}