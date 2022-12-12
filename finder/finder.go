package finder

import (
	"strings"

	ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"
)

type finder struct {
}

func NewFinder() (*finder, error) {
	f := &finder{}

	return f, nil
}

func (f *finder) FindOccurences(text string, sentencesToFind []string) []ti.Word {
	sentecesMap := f.createSentencesMap(sentencesToFind)
	iterator := ti.NewTextIterator(text)

	occurences := []ti.Word{}
	finish := false
	for !finish {
		word := iterator.Next()
		if word == nil {
			finish = true
			continue
		}

		if f.sentenceHasMatch(word.Text, sentecesMap) {
			occurences = append(occurences, *word)
		}
	}

	return occurences
}

func (f *finder) createSentencesMap(sentencesToFind []string) map[string]bool {
	sentencesMap := map[string]bool{}
	for _, sentence := range sentencesToFind {
		lowerCaseSentence := strings.ToLower(sentence)
		sentencesMap[lowerCaseSentence] = true
	}
	return sentencesMap
}

func (f *finder) sentenceHasMatch(sentence string, sentencesMap map[string]bool) bool {
	wordLower := strings.ToLower(sentence)
	return sentencesMap[wordLower]
}
