package finder

import ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"

type Finder interface {
	FindPokemonOccurences(text string) []ti.Word
}
