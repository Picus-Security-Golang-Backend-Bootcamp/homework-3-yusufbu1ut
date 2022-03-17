package author

import (
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) Migration() {
	r.db.AutoMigrate(&Author{})
}

func (r *AuthorRepository) InsertSampleData(authors []Author) {
	for _, c := range authors {
		r.db.Create(&c)
	}
}
