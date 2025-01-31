package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"roman-munteanu/library-kata/models"
	"roman-munteanu/library-kata/repositories"
)

type LibraryApp struct {
	usersRepo repositories.UsersRepositoryAPI
	booksRepo repositories.BooksRepositoryAPI
}

func main() {
	// app init
	a := &LibraryApp{
		usersRepo: repositories.NewUsersRepository(),
		booksRepo: repositories.NewBooksRepository(),
	}

	// handlers
	http.Handle("/users", usersHandler(a))
	http.Handle("/user", userSaveHandler(a))
	http.Handle("/books", booksHandler(a))
	http.Handle("/book", bookSaveHandler(a))
	http.Handle("/borrow", borrowHandler(a))
	http.Handle("/return", returnHandler(a))
	http.Handle("/user-books", userBooksHandler(a))

	// server
	log.Fatalln(http.ListenAndServe(":3000", nil))
}

func usersHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodGet {
			http.Error(rw, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		data, err := app.usersRepo.FetchAll()
		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, data)
	})
}

func userSaveHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			http.Error(rw, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		user := &models.User{}
		err := json.NewDecoder(req.Body).Decode(user)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if user.Name == "" {
			http.Error(rw, "please provide user name", http.StatusBadRequest)
			return
		}

		ok, err := app.usersRepo.Save(*user)
		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, models.Response{IsSuccess: ok})
	})
}

func booksHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodGet {
			http.Error(rw, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		data, err := app.booksRepo.FetchAll()
		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, data)
	})
}

func bookSaveHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			http.Error(rw, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		book := &models.Book{}
		err := json.NewDecoder(req.Body).Decode(book)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if book.Title == "" {
			http.Error(rw, "please provide book title", http.StatusBadRequest)
			return
		}

		ok, err := app.booksRepo.Save(*book)
		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, models.Response{IsSuccess: ok})
	})
}

func borrowHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			http.Error(rw, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		body := &models.Request{}
		err := json.NewDecoder(req.Body).Decode(body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ok, err := app.booksRepo.Borrow(body.UserID, body.BookID)

		var notFoundErr models.NotFoundError
		if errors.As(err, &notFoundErr) {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}

		var genericErr models.GenericError
		if errors.As(err, &genericErr) {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, models.Response{IsSuccess: ok})
	})
}

func returnHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodPost {
			http.Error(rw, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		body := &models.Request{}
		err := json.NewDecoder(req.Body).Decode(body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		ok, err := app.booksRepo.Return(body.UserID, body.BookID)

		var notFoundErr models.NotFoundError
		if errors.As(err, &notFoundErr) {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}

		var genericErr models.GenericError
		if errors.As(err, &genericErr) {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, models.Response{IsSuccess: ok})
	})
}

func userBooksHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if req.Method != http.MethodGet {
			http.Error(rw, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		userID := req.FormValue("user_id")
		if userID == "" {
			http.Error(rw, "provide user_id parameter", http.StatusBadRequest)
			return
		}

		data, err := app.booksRepo.FetchUserBooks(userID)
		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, data)
	})
}

func (app *LibraryApp) toJSONResponse(rw http.ResponseWriter, _ *http.Request, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
	}
}
