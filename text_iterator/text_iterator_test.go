package textiterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNext(t *testing.T) {
	t.Run("test if get the first word", func(t *testing.T) {
		want := &Word{
			Text:    "palavra",
			StartAt: 0,
			EndAt:   6,
		}
		textIterator := NewTextIterator("palavra legal")
		expected := textIterator.Next()
		assert.Equal(t, want, expected, "different word received")
	})

	t.Run("test if get the last word", func(t *testing.T) {
		want := &Word{
			Text:    "legal",
			StartAt: 8,
			EndAt:   12,
		}
		textIterator := NewTextIterator("palavra legal")
		textIterator.Next() // Return palavra
		expected := textIterator.Next()
		require.Equal(t, want, expected, "different word received")
	})

	t.Run("test if get a word in the middle of the text", func(t *testing.T) {
		want := &Word{
			Text:    "diferente",
			StartAt: 8,
			EndAt:   16,
		}
		textIterator := NewTextIterator("palavra diferente legal")
		textIterator.Next() // Return palavra
		expected := textIterator.Next()
		require.Equal(t, want, expected, "different word received when testing the word in the middle")
	})

	t.Run("test if return a nil word", func(t *testing.T) {
		textIterator := NewTextIterator("palavra")
		textIterator.Next() // Return palavra
		expected := textIterator.Next()
		require.Nil(t, expected, "expecting a nil word")
	})

	t.Run("test with a skip char", func(t *testing.T) {
		textIterator := NewTextIterator("\"palavra\"")
		word := textIterator.Next() // Return palavra
		expected := "palavra"
		require.Equal(t, word.Text, expected)
	})

	t.Run("test with farfetch'd", func(t *testing.T) {
		want := &Word{
			Text:    "farfetch'd",
			StartAt: 3,
			EndAt:   12,
		}
		textIterator := NewTextIterator("um farfetch'd")
		textIterator.Next() // Return um
		expected := textIterator.Next()
		require.Equal(t, want, expected, "farfetch'd not recognized")
	})
}
