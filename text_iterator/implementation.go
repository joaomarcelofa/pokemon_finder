package textiterator

type concreteTextIterator struct {
	textInRunes []rune
	index       uint
}

func NewTextIterator(text string) concreteTextIterator {
	return concreteTextIterator{
		textInRunes: []rune(text),
	}
}

func (cti *concreteTextIterator) Next() *Word {
	textHasBeenProcessed := cti.index == uint(len(cti.textInRunes))
	if textHasBeenProcessed {
		return nil
	}

	nextWord := []rune{}
	var word *Word

	i := cti.index
	for {
		char := cti.textInRunes[i]
		wordWasFound := false
		switch {
		case cti.isABreakingRune(char):
			wordWasFound = true
		case cti.isOnLastRune(i):
			nextWord = append(nextWord, char)
			wordWasFound = true
		default:
			nextWord = append(nextWord, char)
		}

		if wordWasFound {
			nextWordStr := string(nextWord)
			nextWordSize := len(nextWord)
			word = &Word{
				Text:    nextWordStr,
				StartAt: cti.index,
				EndAt:   cti.index + uint(nextWordSize-1),
			}
		}

		i += 1
		found := word != nil
		if found {
			cti.index = i
			break
		}
	}

	return word
}

func (cti *concreteTextIterator) isABreakingRune(r rune) bool {
	runeMap := map[rune]bool{
		' ': true,
		',': true,
		'.': true,
		';': true,
	}
	return runeMap[r]
}

func (cti *concreteTextIterator) isOnLastRune(index uint) bool {
	return index == uint(len(cti.textInRunes)-1)
}
