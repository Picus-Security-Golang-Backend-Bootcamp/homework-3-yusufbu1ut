package author

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/pkg/helper"
	"gorm.io/gorm"
)

var id = 0

type Author struct {
	gorm.Model
	ID          uint `gorm:"primarykey"`
	NameSurname string
	Age         int
}

func NewAuthor(nameSurname string) *Author {
	id++
	return &Author{
		ID:          uint(id),
		NameSurname: nameSurname,
		Age:         helper.RandomInt(20, 80),
	}
}

//ToString func for Authors
func (a *Author) ToString() string {
	return fmt.Sprintf("ID : %d, Name Surname: %s, Age: %d", a.ID, a.NameSurname, a.Age)
}

// BeforeDelete for taking info
func (a *Author) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("Author (%s) deleting...", a.NameSurname)
	return nil
}
