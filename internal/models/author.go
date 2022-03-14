package models

type Author struct {
	//Id int
	NameSurname string
}

func NewAuthor(nameSurname string) *Author {
	return &Author{
		NameSurname: nameSurname,
	}
}
