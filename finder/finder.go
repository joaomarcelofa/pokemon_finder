package finder

import (
	"os"
	"strings"

	ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"
)

type finder struct {
	pokemonMap map[string]bool
}

func NewFinder() (*finder, error) {
	listFile := "/home/joao/juit/go/src/github.com/joaomarcelofa/pokemon_finder/pokemon_list.txt"

	fileReader, err := os.Open(listFile)
	if err != nil {
		return nil, err
	}

	pokemonMap := loadList(fileReader)

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
