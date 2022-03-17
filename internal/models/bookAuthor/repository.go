package bookAuthor

import (
	"gorm.io/gorm"
)

type BookAuthRepository struct {
	db *gorm.DB
}

func NewBookAuthRepository(db *gorm.DB) *BookAuthRepository {
	return &BookAuthRepository{
		db: db,
	}
}

func (r *BookAuthRepository) Migration() {
	r.db.AutoMigrate(&Book_Author{})
}
func (r *BookAuthRepository) InsertSampleData(bookAuthhors []Book_Author) {
	for _, c := range bookAuthhors {
		r.db.Create(&c)
	}
}
