package types

// Store is a database interface.
type Store interface {
	FindPeople(textFilter string) ([]Person, error)
}

// Router is the API interface.
type Router interface {
	Get(textFilter string) ([]Person, error)
}
