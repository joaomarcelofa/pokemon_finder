package textiterator

type Word struct {
	Text    string `json:"text"`
	StartAt uint   `json:"start_at"`
	EndAt   uint   `json:"end_at"`
}
