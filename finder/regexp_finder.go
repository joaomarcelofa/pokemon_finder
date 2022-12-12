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
	// textRunes := []rune(textLowerCased)

	pokemons := []ti.Word{}

	for _, pokemon := range sentencesToFind {
		pokemonFormatted := strings.ReplaceAll(pokemon, ".", `\.`)
		pokemonRegex := regexp.MustCompile(fmt.Sprintf(`(?m)%s`, pokemonFormatted))

		occurences := pokemonRegex.FindAllStringIndex(textLowerCased, -1)
		if len(occurences) > 0 {
			for _, occurence := range occurences {
				// https://stackoverflow.com/questions/41956391/how-found-offset-index-a-string-in-rune-using-go
				// https://go.dev/blog/strings
				start := utf8.RuneCountInString(text[:occurence[0]])
				end := utf8.RuneCountInString(text[:occurence[1]]) - 1
				w := ti.Word{
					Text:    text[occurence[0]:occurence[1]],
					StartAt: uint(start),
					EndAt:   uint(end),
				}

				pokemons = append(pokemons, w)
			}
		}
	}

	return pokemons
}
