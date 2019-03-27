package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Age       int      `json:"age,omitempty"`
	Address   *address `json:"address,omitempty"`
}

type address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []person

var errPersonNotFound = errors.New("person not found")

func init() {
	createDummyData()
}

func main() {
	createDummyData()
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}

}

func createDummyData() {
	people = append(people, person{ID: "1", FirstName: "Will", LastName: "Andrews", Age: 29, Address: &address{City: "Here", State: "There"}})

	people = append(people, person{ID: "2", FirstName: "Claire", LastName: "Pask", Age: 30, Address: &address{City: "Here", State: "There"}})

	people = append(people, person{ID: "3", FirstName: "Bea", LastName: "Andrews", Age: 0, Address: &address{City: "Here", State: "There"}})
}

// GetPeople returns all people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson returns a single person
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	person, err := findPerson(params["id"])

	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(person)
}

// CreatePerson creates a new person and adds it to people
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson will delete a person from people
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, person := range people {
		if person.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			json.NewEncoder(w).Encode(people)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

// FindPerson returns a person
func findPerson(id string) (person, error) {

	var p person
	for _, person := range people {
		if person.ID == id {
			p = person
			return p, nil
		}
	}

	return p, errPersonNotFound
}
