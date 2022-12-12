package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joaomarcelofa/pokemon_finder/finder"
	ll "github.com/joaomarcelofa/pokemon_finder/list_loader"
)

type pokemonHandler struct {
	finder      finder.TextFinder
	pokemonList []string
}

func newPokemonHandler() (*pokemonHandler, error) {
	finder, err := finder.NewFinder()
	if err != nil {
		return nil, err
	}

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	paths := strings.SplitAfter(dir, "pokemon_finder")
	if len(paths) > 1 {
		dir = paths[0]
	}

	listFile := fmt.Sprintf("%s/pokemon_list.txt", dir)
	fileReader, err := os.Open(listFile)
	if err != nil {
		return nil, err
	}

	pokemonList := ll.LoadList(fileReader)
	pokemonHandler := &pokemonHandler{
		finder:      finder,
		pokemonList: pokemonList,
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

	pokemons := ph.finder.FindOccurences(textBody.Text, ph.pokemonList)
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
