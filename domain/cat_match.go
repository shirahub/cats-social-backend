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
