package sample

import (
	"encoding/csv"
	"strconv"

	"os"
	"strings"
	"sync"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/author"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models/bookAuthor"
)

var ResultsBooks []book.Book
var ResultsAuthors []author.Author
var ResultsBookAuth []bookAuthor.Book_Author

type item struct {
	bks   book.Book
	auths string
}

func ReadBookWithWorkerPool(path string) error {

	linesChan := make(chan []string)
	resultsChan := make(chan item)
	var results []item
	wgB := new(sync.WaitGroup)
	for w := 1; w <= 3; w++ {
		wgB.Add(1)
		go convertToItemsStruct(linesChan, resultsChan, wgB)
	}

	go func() {
		f, _ := os.Open(path)
		defer f.Close()
		lines, _ := csv.NewReader(f).ReadAll()
		isFirstRow := true
		for _, line := range lines {
			if isFirstRow {
				isFirstRow = false
				continue
			}

			linesChan <- line
		}
		close(linesChan)
	}()

	go func() {
		wgB.Wait()

		close(resultsChan)
	}()

	for i := range resultsChan {
		results = append(results, i)
	}
	for _, v := range results {
		ResultsBooks = append(ResultsBooks, v.bks)
		sepAddAuthors(v.auths, v.bks)
	}
	// fmt.Println(len(ResultsBooks))
	// fmt.Println(len(ResultsAuthors))
	// fmt.Println(len(ResultsBookAuth))

	return nil
}

func convertToItemsStruct(booksChan <-chan []string, resultschan chan<- item, wg *sync.WaitGroup) {
	defer wg.Done()

	for b := range booksChan {
		isbn, _ := strconv.Atoi(b[5])
		_book := book.NewBook(b[1], isbn, b[7])
		_item := item{
			bks:   *_book,
			auths: b[2],
		}
		resultschan <- _item
	}
}

func sepAddAuthors(auths string, book book.Book) {
	authsRes := strings.Split(auths, "/")
	for _, a := range authsRes {
		_author := isInAuthors(a)
		_bookAuthor := bookAuthor.NewBook_Author(book.ISBN, _author.ID)
		ResultsBookAuth = append(ResultsBookAuth, *_bookAuthor)
	}
}

func isInAuthors(authr string) author.Author {
	for _, v := range ResultsAuthors {
		if v.NameSurname == authr {
			return v
		}
	}
	_author := author.NewAuthor(authr)
	ResultsAuthors = append(ResultsAuthors, *_author)
	return *_author

}
