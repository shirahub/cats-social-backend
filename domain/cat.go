package domain

type Cat struct {
	Id          string
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	ImageUrls   []string
	UserId      string
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
