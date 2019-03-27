package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetPerson(t *testing.T) {

	t.Run("return correct person", func(t *testing.T) {
		want := person{ID: "1"}

		got, err := findPerson("1")

		AssertError(t, nil, err)

		if got.ID != want.ID {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("person not found", func(t *testing.T) {
		want := person{}

		got, err := findPerson("99")

		AssertError(t, errPersonNotFound, err)

		if got.ID != want.ID {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func TestGetPeopleHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/people", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPeople)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	got := rr.Body.String()
	want := `[{"id":"1","firstname":"Will","lastname":"Andrews","age":29,"address":{"city":"Here","state":"There"}},{"id":"2","firstname":"Claire","lastname":"Pask","age":30,"address":{"city":"Here","state":"There"}},{"id":"3","firstname":"Bea","lastname":"Andrews","address":{"city":"Here","state":"There"}}]`

	if strings.TrimSpace(got) != strings.TrimSpace(want) {
		t.Errorf("got %s but want %s", got, want)
	}

}

func TestGetPersonHander(t *testing.T) {
	req, err := http.NewRequest("GET", "/people/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")

	r.ServeHTTP(rr, req)

	want := `{"id":"1","firstname":"Will","lastname":"Andrews","age":29,"address":{"city":"Here","state":"There"}}`

	got := rr.Body.String()

	if strings.TrimSpace(got) != strings.TrimSpace(want) {
		t.Errorf("got %s but want %s", got, want)
	}

}

func AssertError(t *testing.T, want, got error) {
	if got != want {
		t.Errorf("wanted error %v but got %v", want, got)
	}
}
