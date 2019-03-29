package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/willdot/Go-HttpServerDemo/store"

	"github.com/gorilla/mux"
)

var people []store.Person

func createDummyData() {
	people = append(people, store.Person{ID: "1", FirstName: "Will", LastName: "Andrews", Age: 29, Address: &store.Address{City: "Here", State: "There"}})

	people = append(people, store.Person{ID: "2", FirstName: "Claire", LastName: "Pask", Age: 30, Address: &store.Address{City: "Here", State: "There"}})

	people = append(people, store.Person{ID: "3", FirstName: "Bea", LastName: "Andrews", Age: 0, Address: &store.Address{City: "Here", State: "There"}})
}

type FakeStore struct {
	returnError bool
}

func (f *FakeStore) Create(p *store.Person) error {
	if f.returnError {
		return errPersonAlreadyExists
	}
	return nil
}

func (f *FakeStore) Delete(ID string) error {
	if f.returnError {
		return errPersonNotFound
	}
	return nil
}

func (f *FakeStore) FindByID(ID string) (store.Person, error) {
	var person store.Person
	if f.returnError {
		return person, errPersonNotFound
	}
	person = store.Person{ID: "1", FirstName: "Will", LastName: "Andrews", Age: 29, Address: &store.Address{City: "Here", State: "There"}}
	return person, nil
}

func (f *FakeStore) GetAll() ([]store.Person, error) {
	if f.returnError {
		return nil, errors.New("Can't find anyone")
	}
	createDummyData()
	return people, nil
}

func getFakeApp(returnError bool) store.App {
	fakeDb := new(FakeStore)
	fakeDb.returnError = returnError

	fakeApp := store.App{DB: fakeDb}
	return fakeApp
}

func TestGetPeopleHandler(t *testing.T) {

	t.Run("Returns people", func(t *testing.T) {
		fakeApp := getFakeApp(false)
		req, err := http.NewRequest("GET", "/people", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetPeople(&fakeApp))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		got := rr.Body.String()
		want := `[{"id":"1","firstname":"Will","lastname":"Andrews","age":29,"address":{"city":"Here","state":"There"}},{"id":"2","firstname":"Claire","lastname":"Pask","age":30,"address":{"city":"Here","state":"There"}},{"id":"3","firstname":"Bea","lastname":"Andrews","address":{"city":"Here","state":"There"}}]`

		if strings.TrimSpace(got) != strings.TrimSpace(want) {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("Returns error", func(t *testing.T) {
		fakeApp := getFakeApp(true)
		req, err := http.NewRequest("GET", "/people", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetPeople(&fakeApp))

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestGetPersonHander(t *testing.T) {

	t.Run("Returns a person", func(t *testing.T) {
		fakeApp := getFakeApp(false)
		req, err := http.NewRequest("GET", "/people/1", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()

		r.HandleFunc("/people/{id}", GetPerson(&fakeApp)).Methods("GET")

		r.ServeHTTP(rr, req)

		want := `{"id":"1","firstname":"Will","lastname":"Andrews","age":29,"address":{"city":"Here","state":"There"}}`

		got := rr.Body.String()

		if strings.TrimSpace(got) != strings.TrimSpace(want) {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("Returns not found", func(t *testing.T) {
		fakeApp := getFakeApp(true)
		req, err := http.NewRequest("GET", "/people/1", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()

		r.HandleFunc("/people/{id}", GetPerson(&fakeApp)).Methods("GET")

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})

}

func TestDeletePerson(t *testing.T) {
	t.Run("Returns 200 ok", func(t *testing.T) {
		fakeApp := getFakeApp(false)
		req, err := http.NewRequest("DELETE", "/people/1", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/people/{id}", DeletePerson(&fakeApp)).Methods("DELETE")

		r.ServeHTTP(rr, req)

		want := http.StatusOK
		got := rr.Code

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}
	})

	t.Run("Returns 404 not found", func(t *testing.T) {
		fakeApp := getFakeApp(true)
		req, err := http.NewRequest("DELETE", "/people/1", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/people/{id}", DeletePerson(&fakeApp)).Methods("DELETE")

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})

}

func TestCreatePerson(t *testing.T) {
	t.Run("Returns people", func(t *testing.T) {
		fakeApp := getFakeApp(false)
		req, err := http.NewRequest("POST", "/people/99", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/people/{id}", CreatePerson(&fakeApp)).Methods("POST")

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})

	t.Run("Returns error", func(t *testing.T) {
		fakeApp := getFakeApp(true)
		req, err := http.NewRequest("POST", "/people/99", nil)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r := mux.NewRouter()
		r.HandleFunc("/people/{id}", CreatePerson(&fakeApp)).Methods("POST")

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

		errMessage := strings.TrimSpace(rr.Body.String())

		if errMessage != errPersonAlreadyExists.Error() {
			t.Errorf("got %s but want %s", errMessage, errPersonAlreadyExists.Error())
		}

	})
}

func AssertError(t *testing.T, want, got error) {
	if got != want {
		t.Errorf("wanted error %v but got %v", want, got)
	}
}
