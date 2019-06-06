package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/frankienicoletti/fullstack-exercise/src/server/types"
	"github.com/gorilla/mux"
)

var _ types.Router = &PeopleAPI{}

// PeopleAPI represents an instance of API.
type PeopleAPI struct {
	db     types.Store
	Router *mux.Router
}

// NewPeopleAPI returns a new instance of an PeopleAPI.
func NewPeopleAPI(db types.Store) *PeopleAPI {
	peopleAPI := &PeopleAPI{db: db}

	router := mux.NewRouter()
	router.HandleFunc("/people", peopleAPI.get).Methods("GET")
	peopleAPI.Router = router

	return peopleAPI
}

// Get retrieves people given optional name filter.
func (p *PeopleAPI) Get(textFilter string) ([]types.Person, error) {
	return p.db.FindPeople(textFilter)
}

func (p *PeopleAPI) get(w http.ResponseWriter, req *http.Request) {
	textFilter := req.FormValue("q")

	people, err := p.Get(textFilter)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	} else if len(people) == 0 {
		w.Write([]byte(`[]`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(people)
}
