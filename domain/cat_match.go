package domain

type CatMatch struct {
	Id            string
	Message       string
	IssuerCatId   string `json:"userCatDetail"`
	ReceiverCatId string `json:"matchCatDetail"`
	Status        string
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:",omitempty"`
}
