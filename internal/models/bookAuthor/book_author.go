package bookAuthor

import "gorm.io/gorm"

type Book_Author struct {
	gorm.Model
	BookID   int  `gorm:"foreingkey"`
	AuthorID uint `gorm:"foreingkey"`
}

func NewBook_Author(bookisbn int, authorid uint) *Book_Author {
	return &Book_Author{
		BookID:   bookisbn,
		AuthorID: authorid,
	}
}
