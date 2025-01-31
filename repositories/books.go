package repositories

import (
	"roman-munteanu/library-kata/models"
)

var booksData map[string]models.Book

type BooksRepositoryAPI interface {
	FetchAll() ([]models.Book, error)
	Borrow(userID, bookID string) (bool, error)
	Return(userID, bookID string) (bool, error)
}

type BooksRepository struct {
}

func NewBooksRepository() *BooksRepository {

	booksData = map[string]models.Book{
		"d9914b4f-d0ec-405b-b1cc-1387d489bc5e": {
			ID:          "d9914b4f-d0ec-405b-b1cc-1387d489bc5e",
			Title:       "Book 1",
			TakenByUser: "",
		},
		"55f6d196-b0cf-46c6-8aba-dfb8315249d4": {
			ID:          "55f6d196-b0cf-46c6-8aba-dfb8315249d4",
			Title:       "Book 2",
			TakenByUser: "",
		},
		"cf99c3dd-08b9-4b29-86eb-d981219061fb": {
			ID:          "cf99c3dd-08b9-4b29-86eb-d981219061fb",
			Title:       "Book 3",
			TakenByUser: "",
		},
	}

	return &BooksRepository{}
}

func (r *BooksRepository) FetchAll() ([]models.Book, error) {
	var books []models.Book
	for _, v := range booksData {
		books = append(books, v)
	}
	return books, nil
}

func (r *BooksRepository) Borrow(userID, bookID string) (bool, error) {
	book, ok := booksData[bookID]
	if !ok {
		return false, nil
	}

	// already taken
	if book.TakenByUser != "" {
		return false, nil
	}

	book.TakenByUser = userID
	booksData[bookID] = book

	return true, nil
}

func (r *BooksRepository) Return(userID, bookID string) (bool, error) {
	book, ok := booksData[bookID]
	if !ok {
		return false, nil
	}

	// not taken by current user
	if book.TakenByUser != userID {
		return false, nil
	}

	book.TakenByUser = ""
	booksData[bookID] = book

	return true, nil
}
