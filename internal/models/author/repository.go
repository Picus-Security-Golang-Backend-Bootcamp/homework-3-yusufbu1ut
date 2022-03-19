package author

import (
	"errors"

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

// InsertSampleData adds datas from csv when program runs
func (r *AuthorRepository) InsertSampleData(authors []Author) {
	for _, c := range authors {
		r.db.Create(&c)
	}
}

// GetAllAuthors takes all authors
func (r *AuthorRepository) GetAllAuthors() ([]Author, error) {
	var authors []Author
	result := r.db.Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

//FindByAuthorName finds all Authors with given input
func (r *AuthorRepository) FindByAuthorName(input string) ([]Author, error) {
	var authors []Author
	result := r.db.Where("LOWER(name_surname) LIKE LOWER(?)", "%"+input+"%").Find(&authors)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

//FindByID finds author with given id
func (r *AuthorRepository) FindByID(id int) ([]Author, error) {
	var authors []Author
	result := r.db.Find(&authors, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return authors, nil
}

//Delete deletes author from base with parameter author
func (r *AuthorRepository) Delete(a Author) error {
	result := r.db.Delete(a)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//DeleteByID deletes author from base with given id
func (r *AuthorRepository) DeleteByID(id int) error {
	result := r.db.Delete(&Author{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

//Create creats with given author
func (r *AuthorRepository) Create(author Author) error {
	var athr Author
	result := r.db.Where(&Author{NameSurname: author.NameSurname, Age: author.Age}).First(&athr)
	if result.Error != nil {
		return result.Error
	}
	if athr.NameSurname == author.NameSurname {
		return errors.New("Author has been declared")
	}
	r.db.Create(author)
	return nil
}

//Update updates with given author
func (r *AuthorRepository) Update(author Author) error {
	r.db.Save(author)
	return nil
}
