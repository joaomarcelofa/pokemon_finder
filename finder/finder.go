package finder

import (
	"fmt"
	"os"
	"strings"

	ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"
)

type finder struct {
	pokemonMap map[string]bool
}

func NewFinder() (*finder, error) {
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

	pokemonMap := loadMap(fileReader)

	f := &finder{
		pokemonMap: pokemonMap,
	}

	return f, nil
}

func (f *finder) FindPokemonOccurences(text string) []ti.Word {
	iterator := ti.NewTextIterator(text)

	pokemons := []ti.Word{}
	finish := false
	for !finish {
		word := iterator.Next()
		if word == nil {
			finish = true
			continue
		}

		if f.isPokemon(word.Text) {
			pokemons = append(pokemons, *word)
		}
	}

	return pokemons
}

func (f *finder) isPokemon(word string) bool {
	wordLower := strings.ToLower(word)
	return f.pokemonMap[wordLower]
}
