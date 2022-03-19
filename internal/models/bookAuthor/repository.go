package bookAuthor

import (
	"errors"

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

//InsertSampleData creates datas for book_author
func (r *BookAuthRepository) InsertSampleData(bookAuthhors []Book_Author) {
	for _, c := range bookAuthhors {
		r.db.Create(&c)
	}
}

// FindByISBN finds elements with isbn count it is for book connections
func (r *BookAuthRepository) FindByISBN(isbn int) ([]Book_Author, error) {
	var book_authors []Book_Author
	result := r.db.Find(&book_authors, "book_id = ?", isbn)

	if result.Error != nil {
		return nil, result.Error
	}

	return book_authors, nil
}

//FindByAuthorID finds elements with using author id
func (r *BookAuthRepository) FindByAuthorID(id int) ([]Book_Author, error) {
	var book_authors []Book_Author
	result := r.db.Find(&book_authors, "author_id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return book_authors, nil
}

//DeleteByISBN deletes items looking isbn colmn returns authors that have no book which as not deleted
func (r *BookAuthRepository) DeleteByISBN(isbn int) ([]int, error) {
	ba, _ := r.FindByISBN(isbn)
	result := r.db.Delete(ba)
	var authorsID []int
	for _, v := range ba {
		var book_authors []Book_Author
		rslt := r.db.Find(&book_authors, "author_id = ?", v.AuthorID)
		if rslt.Error != nil {
			return nil, rslt.Error
		}
		if len(book_authors) == 0 {
			authorsID = append(authorsID, int(v.AuthorID))
		}
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return authorsID, nil
}

// Create creates new book_authors element
func (r *BookAuthRepository) Create(bauthors Book_Author) error {
	var ba Book_Author
	result := r.db.Where(&Book_Author{BookID: bauthors.BookID, AuthorID: bauthors.AuthorID}).First(&ba)
	if result.Error != nil {
		return result.Error
	}
	if ba.AuthorID == bauthors.AuthorID {
		return errors.New("Element has been declared")
	}
	r.db.Create(bauthors)
	return nil
}
