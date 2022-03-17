package main

import (
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/infrastructure"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/bookAuthor"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/sample"
)

var (
	bookRepository       *book.BookRepository
	authorRepository     *author.AuthorRepository
	bookAuthorRepository *bookAuthor.BookAuthRepository
)

func init() {
	// //db := infrastructure.NewMySQLDB("root:Password123!@tcp(127.0.0.1:3306)/location?parseTime=True&loc=Local")
	sample.ReadBookWithWorkerPool("../docs/books.csv")
	db := infrastructure.NewPostgresDB("host=localhost user=postgres password=pass1234 dbname=library port=5432 sslmode=disable")
	bookRepository = book.NewBookRepository(db)
	authorRepository = author.NewAuthorRepository(db)
	bookAuthorRepository = bookAuthor.NewBookAuthRepository(db)
	bookRepository.Migration()
	bookRepository.InsertSampleData(sample.ResultsBooks)
	authorRepository.Migration()
	authorRepository.InsertSampleData(sample.ResultsAuthors)
	bookAuthorRepository.Migration()
	bookAuthorRepository.InsertSampleData(sample.ResultsBookAuth)

}

func main() {
	// news := author.NewAuthor("asd")
	// fmt.Println(news)

	// sample.ReadBookWithWorkerPool("../docs/books.csv")

	// for i, _ := range sample.ResultsAuthors {

	// 	for j := i + 1; j < len(sample.ResultsAuthors); j++ {
	// 		if sample.ResultsAuthors[i] == sample.ResultsAuthors[j] {
	// 			fmt.Println(sample.ResultsAuthors[i])
	// 			break
	// 		}
	// 	}

	// }
	// fmt.Println(len(sample.ResultsBookAuth))
	// fmt.Println(len(sample.ResultsAuthors))
	fmt.Scanln()
}
