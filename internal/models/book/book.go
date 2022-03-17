package book

import (
	"gorm.io/gorm"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/pkg/helper"
)

type Book struct {
	gorm.Model  //`gorm:"foreignKey:BookID;references:ID"`
	Name        string
	StockCode   int //random
	ISBN        int
	Pages       string
	Price       float64 //random
	StockAmount int     //ramdom

}

func NewBook(name string, isbn int, page string) *Book {
	return &Book{
		Name:        name,
		ISBN:        isbn,
		Pages:       page,
		StockCode:   helper.RandomInt(100000, 1000000),
		Price:       helper.RandFloat(50, 250),
		StockAmount: helper.RandomInt(50, 100),
	}
}
