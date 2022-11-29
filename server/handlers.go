package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joaomarcelofa/pokemon_finder/finder"
)

type pokemonHandler struct {
	finder finder.Finder
}

func newPokemonHandler() (*pokemonHandler, error) {
	finder, err := finder.NewFinder()
	if err != nil {
		return nil, err
	}
	pokemonHandler := &pokemonHandler{
		finder: finder,
	}

	return pokemonHandler, nil
}

func (ph *pokemonHandler) get(w http.ResponseWriter, req *http.Request) {
	log.Println("received a request")

	decoder := json.NewDecoder(req.Body)
	var textBody input
	err := decoder.Decode(&textBody)
	if err != nil {
		log.Fatalf("error at decoding body")
		return
	}

	pokemons := ph.finder.FindPokemonOccurences(textBody.Text)
	log.Println("request processed")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(pokemons)
	if err != nil {
		log.Fatalf("error at forming the response")
		return
	}

	log.Println("returning the processed request")
	w.Write(jsonResponse)
}
