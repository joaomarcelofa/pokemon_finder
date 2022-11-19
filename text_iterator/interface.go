package textiterator

type TextIterator interface {
	// Next Return the next word
	Next() *Word
}
