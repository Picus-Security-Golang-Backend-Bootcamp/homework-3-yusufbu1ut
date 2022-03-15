package models

import (
	//"sync"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/pkg/helper"
)

type Book struct {
	//sync.RWMutex
	ID          int     `json:"id"`
	Name        string  `json:"bookName"`
	Author      Author  `json:"authorName"`
	StockCode   int     `json:"stockCode"` //random`
	ISBN        int     `json:"isbn"`
	Pages       int     `json:"pages"`
	Price       float64 `json:"price"`       //random
	StockAmount int     `json:"stockAmount"` //ramdom
}

func NewBook(id int, name string, auth Author, isbn int, page int) *Book {
	return &Book{
		ID:          id,
		Name:        name,
		Author:      auth,
		ISBN:        isbn,
		Pages:       page,
		StockCode:   helper.RandomInt(100000, 1000000),
		Price:       helper.RandFloat(50, 250),
		StockAmount: helper.RandomInt(50, 100),
	}
}
