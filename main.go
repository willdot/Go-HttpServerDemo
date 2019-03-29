//+build !test

package main

import (
	"Go-HttpServerDemo/handlers"
	"Go-HttpServerDemo/store"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var app store.App

func init() {

	db := new(store.RealStore)

	app = store.App{DB: db}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/people", handlers.GetPeople(&app)).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPerson(&app)).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.CreatePerson(&app)).Methods("POST")
	router.HandleFunc("/people/{id}", handlers.DeletePerson(&app)).Methods("DELETE")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}

}
