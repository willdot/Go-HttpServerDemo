package store

type (
	// Store is an interface for data storage
	Store interface {
		Create(p *Person) error
		Delete(ID string) error
		FindByID(ID string) (Person, error)
		GetAll() ([]Person, error)
	}

	// App is a struct that contains a DB store
	App struct {
		DB Store
	}

	// Person is a person model
	Person struct {
		ID        string   `json:"id,omitempty"`
		FirstName string   `json:"firstname,omitempty"`
		LastName  string   `json:"lastname,omitempty"`
		Age       int      `json:"age,omitempty"`
		Address   *Address `json:"address,omitempty"`
	}

	// Address is an address model
	Address struct {
		City  string `json:"city,omitempty"`
		State string `json:"state,omitempty"`
	}
)

// RealStore is a proper implementation of a store
type RealStore struct {
}

// Create will create a new person
func (f *RealStore) Create(p *Person) error {
	return nil
}

// Delete will delete a person by id
func (f *RealStore) Delete(ID string) error {
	return nil
}

// FindByID returns a person using an id
func (f *RealStore) FindByID(ID string) (Person, error) {
	person := Person{}
	return person, nil
}

//GetAll returns all people
func (f *RealStore) GetAll() ([]Person, error) {
	return nil, nil
}
