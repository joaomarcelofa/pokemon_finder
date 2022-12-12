package finder

import (
	"log"
	"reflect"
	"testing"

	ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newFinder() *finder {
	finder, err := NewFinder()
	if err != nil {
		log.Fatalf("can't create the finder: %s", err)
	}
	return finder
}

func BenchmarkFindOccurences(b *testing.B) {
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

	f := newFinder()
	found := f.FindOccurences(text, pokemonList)

	if !reflect.DeepEqual(found, want) {
		b.Errorf("want: %v, received: %v", want, found)
	}
}

func TestNewFinder(t *testing.T) {
	finder, err := NewFinder()
	require.NoError(t, err)
	assert.NotNil(t, finder)
}
