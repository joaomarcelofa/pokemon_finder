package main

import (
	"log"

	"github.com/joaomarcelofa/pokemon_finder/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		log.Fatalln("error at creating server")
	}

	s.Listen()
}
