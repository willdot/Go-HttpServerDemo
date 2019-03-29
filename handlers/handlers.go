package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"Go-HttpServerDemo/store"

	"github.com/gorilla/mux"
)

var errPersonNotFound = errors.New("person not found")
var errPersonAlreadyExists = errors.New("This person can't be created")

// GetPeople returns all people
func GetPeople(a *store.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		people, err := a.DB.GetAll()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(people)
	}
}

// GetPerson returns a single person
func GetPerson(a *store.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		person, err := a.DB.FindByID(params["id"])

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(person)
	}
}

// CreatePerson creates a new person and adds it to people
func CreatePerson(a *store.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		var person store.Person
		person.ID = params["id"]
		err := a.DB.Create(&person)

		if err != nil {
			http.Error(w, string(errPersonAlreadyExists.Error()), http.StatusBadRequest)
			return
		}

		people, _ := a.DB.GetAll()
		json.NewEncoder(w).Encode(people)
	}

}

// DeletePerson will delete a person from people
func DeletePerson(a *store.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		err := a.DB.Delete(params["id"])

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

		return
	}
}
