package main

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/repository"
)

func init() {

}

func main() {
	a := models.NewAuthor("b")
	bo := models.NewBook(1, "a", *a, 1, 1)
	fmt.Println(bo.ID)

	fmt.Println("")

	repository.ReadBookWithWorkerPool("../docs/books.csv")
	fmt.Scanln()
}
