package finder

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	ti "github.com/joaomarcelofa/pokemon_finder/text_iterator"
)

type regexpFinder struct {
}

func NewRegexpFinder() (*regexpFinder, error) {
	rf := &regexpFinder{}

	return rf, nil
}

func (rf *regexpFinder) FindOccurences(text string, sentencesToFind []string) []ti.Word {
	textLowerCased := strings.ToLower(text)

	occurences := []ti.Word{}

	for _, sentence := range sentencesToFind {
		sentenceFormatted := strings.ReplaceAll(sentence, ".", `\.`)
		sentenceRegex := regexp.MustCompile(fmt.Sprintf(`(?m)%s`, sentenceFormatted))

		occurenceIndexes := sentenceRegex.FindAllStringIndex(textLowerCased, -1)
		if len(occurenceIndexes) > 0 {
			for _, occurenceIndex := range occurenceIndexes {
				// https://stackoverflow.com/questions/41956391/how-found-offset-index-a-string-in-rune-using-go
				// https://go.dev/blog/strings
				start := utf8.RuneCountInString(text[:occurenceIndex[0]])
				end := utf8.RuneCountInString(text[:occurenceIndex[1]]) - 1
				w := ti.Word{
					Text:    text[occurenceIndex[0]:occurenceIndex[1]],
					StartAt: uint(start),
					EndAt:   uint(end),
				}

				occurences = append(occurences, w)
			}
		}
	}

	return occurences
}
