package main

import (
	"encoding/json"
	"log"
	"net/http"

	"roman-munteanu/library-kata/repositories"
)

type LibraryApp struct {
	usersRepo repositories.UsersRepositoryAPI
}

func main() {
	// app init
	a := &LibraryApp{
		usersRepo: repositories.NewUsersRepository(),
	}

	// handlers
	http.Handle("/", usersHandler(a))

	// server
	log.Fatalln(http.ListenAndServe(":3000", nil))
}

func usersHandler(app *LibraryApp) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		users, err := app.usersRepo.FetchAll()
		if err != nil {
			http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		app.toJSONResponse(rw, req, users)
	})
}

func (app *LibraryApp) toJSONResponse(rw http.ResponseWriter, _ *http.Request, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		http.Error(rw, http.StatusText(500), http.StatusInternalServerError)
	}
}
