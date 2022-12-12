package finder

import ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"

type TextFinder interface {
	FindOccurences(text string, sentencesToFind []string) []ti.Word
}
