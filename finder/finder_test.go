package finder

import (
	"reflect"
	"testing"

	ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"
)

func BenchmarkFindPokemonOccurences(b *testing.B) {
	text := "Eu sou um treinador pokémon e tenho um Pikachu, mas também tenho um Gengar"

	want := []ti.Word{
		{
			Text:    "Pikachu",
			StartAt: 39,
			EndAt:   45,
		},
		{
			Text:    "Gengar",
			StartAt: 68,
			EndAt:   73,
		},
	}

	found := FindPokemonOccurences(text)

	if !reflect.DeepEqual(found, want) {
		b.Errorf("want: %v, received: %v", want, found)
	}
}
