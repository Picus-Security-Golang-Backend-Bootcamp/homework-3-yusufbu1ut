package models

type Author struct {
	//Id int
	NameSurname string `json:"AuthorName"`
}

func NewAuthor(nameSurname string) *Author {
	return &Author{
		NameSurname: nameSurname,
	}
}
