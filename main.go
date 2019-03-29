//+build !test

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/willdot/Go-HttpServerDemo/handlers"
	"github.com/willdot/Go-HttpServerDemo/store"

	"github.com/gorilla/mux"
)

var app store.App

func init() {

	db := new(store.RealStore)

	app = store.App{DB: db}
}

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	router := mux.NewRouter()

	router.HandleFunc("/people", handlers.GetPeople(&app)).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPerson(&app)).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.CreatePerson(&app)).Methods("POST")
	router.HandleFunc("/people/{id}", handlers.DeletePerson(&app)).Methods("DELETE")
	err := http.ListenAndServe(":"+PORT, router)

	if err != nil {
		log.Fatal(err)
	}

}
