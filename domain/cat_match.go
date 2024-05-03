package domain

import "time"

type CatMatch struct {
	Id            string
	Message       string
	IssuerCatId   string    `json:"userCatDetail"`
	ReceiverCatId string    `json:"matchCatDetail"`
	Status        string
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:",omitempty"`
}

type CatMatchDetail struct {
	Id string
	Message string
	Status string
	CatMatchCreatedAt time.Time
	IssuerCat Cat
	ReceiverCat Cat
	UserId string
	Name string
	Email string
	UserCreatedAt time.Time
}
