package finder

import (
	"reflect"
	"testing"
)

func BenchmarkFindPokemon(b *testing.B) {
	text := "Eu sou um treinador pokémon e tenho um Pikachu, mas também tenho um Gengar"

	want := []Occurence{
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

	found := FindPokemon(text)

	if !reflect.DeepEqual(found, want) {
		b.Errorf("want: %v, received: %v", want, found)
	}
}
