package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-yusufbu1ut/internal/models"
)

func ReadBookWithWorkerPool(path string) error {

	books := make(chan []string)
	results := make(chan models.Book)

	wg := new(sync.WaitGroup)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go convertToBookStruct(books, results, wg)
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

			books <- line
		}
		close(books)
	}()

	go func() {
		wg.Wait()

		close(results)
	}()

	for v := range results {
		fmt.Println(v)
	}

	return nil
}

func convertToBookStruct(books <-chan []string, results chan<- models.Book, wg *sync.WaitGroup) {
	defer wg.Done()

	for b := range books {
		id, _ := strconv.Atoi(b[0]) //hatalar çıktılana bilinir
		isbn, _ := strconv.Atoi(b[4])
		page, _ := strconv.Atoi(b[7])
		auth := models.NewAuthor(b[2])

		book := models.NewBook(id, b[1], *auth, isbn, page)

		results <- *book
	}
}
