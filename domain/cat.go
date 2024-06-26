package domain

import "time"

type Cat struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  int       `json:"ageInMonth"`
	Description string    `json:"description"`
	ImageUrls   []string  `json:"imageUrls"`
	UserId      string    `json:"userId"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"-"`
}

type CreateCatRequest struct {
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	ImageUrls   []string
	UserId      string
}

type GetCatsRequest struct {
	Id         string
	Limit      int
	Offset     int
	Race       string
	Sex        string
	HasMatched *bool
	AgeInMonth string
	UserId     string
	Owned      *bool
	Name       string
}
