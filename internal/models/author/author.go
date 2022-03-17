package author

import "gorm.io/gorm"

var id = 0

type Author struct {
	gorm.Model
	ID          uint `gorm:"primarykey"`
	NameSurname string
}

func NewAuthor(nameSurname string) *Author {
	id++
	return &Author{
		ID:          uint(id),
		NameSurname: nameSurname,
	}
}
