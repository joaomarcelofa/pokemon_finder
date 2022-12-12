package finder

import (
	"log"
	"reflect"
	"testing"

	ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"
)

func getRegexpFinder() *regexpFinder {
	rf, err := NewRegexpFinder()
	if err != nil {
		log.Fatalf("can't create the regex finder: %s", err)
	}
	return rf
}

func BenchmarkFindOccurencesRegexp(b *testing.B) {
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

	f := getRegexpFinder()
	found := f.FindOccurences(text, pokemonList)

	if !reflect.DeepEqual(found, want) {
		b.Errorf("want: %v, received: %v", want, found)
	}
}
