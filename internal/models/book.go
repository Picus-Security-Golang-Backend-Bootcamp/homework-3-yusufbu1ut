package models

import (
	"sync"
)

type Book struct {
	sync.RWMutex
	ID        int
	Name      string
	Author    Author
	StockCode string //random
	ISBN      int
	Pages     int
	Price     float64 //random
	Amount    int     //ramdom
}

func NewBook(id int, name string, auth Author, isbn int, page int) *Book {
	return &Book{
		ID:     id,
		Name:   name,
		Author: auth,
		ISBN:   isbn,
		Pages:  page,
	}
}
