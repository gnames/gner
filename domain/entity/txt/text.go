package txt

import "github.com/gnames/gnlib/encode"

type textNER struct {
	Lines    map[int]int `json:"lines"`
	Text     []rune
	Entities []EntityNER `json:"entities"`
}

func NewTextNER(text []rune) TextNER {
	return &textNER{Text: text}
}

func (tn *textNER) SetLines(lines map[int]int) {
	tn.Lines = lines
}

func (tn *textNER) GetLines() map[int]int {
	return tn.Lines
}

func (tn *textNER) GetText() []rune {
	return tn.Text
}

func (tn *textNER) SetEntities(ents []EntityNER) {
	tn.Entities = ents
}

func (tn *textNER) GetEntities() []EntityNER {
	return tn.Entities
}

func (tn *textNER) ToJSON(pretty bool) ([]byte, error) {
	enc := encode.GNjson{Pretty: pretty}
	return enc.Encode(tn)
}
