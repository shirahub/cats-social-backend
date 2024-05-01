package domain

type Cat struct {
	Id          uint
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	ImageUrls   []string
}

type CreateCatRequest struct {
	Name        string
	Race        string
	Sex         string
	AgeInMonth  int
	Description string
	ImageUrls   []string
	UserId      uint
}
