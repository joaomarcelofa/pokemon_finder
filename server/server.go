package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	pokemonHandler *pokemonHandler
}

func NewServer() (*server, error) {
	pokeHandler, err := newPokemonHandler()
	if err != nil {
		return nil, err
	}
	s := &server{
		pokemonHandler: pokeHandler,
	}
	return s, nil
}

func (s *server) Listen() {
	handler := s.mountHandler()
	log.Println("server started")
	http.ListenAndServe(":7500", handler)
}

func (s *server) mountHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/pokemon", s.pokemonHandler.get).Methods("GET")
	return router
}
