package txt

import "github.com/gnames/gnfmt"

type textNER struct {
	// Text is the text of the input. It is the content that will be used for
	// named entity recognition.
	Text []rune

	// Entities is a list found in text entities. They can be scientific names,
	// numbers, geographical names etc.
	Entities []EntityNER `json:"entities"`

	// LineEntitiesNum holds the number of entities found on each line
	// the text. The key is the line number in the text, the value is the number
	// of entities found at the line.
	LineEntitiesNum map[int]int `json:"lines"`
}

// NewTextNER takes the content of a text and returns a TextNER compatible
// object.
func NewTextNER(text []rune) TextNER {
	return &textNER{Text: text}
}

// SetLines saves the detected lines of text into TextNER object.
func (tn *textNER) SetLinesEntitiesNum(lines map[int]int) {
	tn.LineEntitiesNum = lines
}

// GetLines retrieves stored lines of text
func (tn *textNER) GetLinesEntitiesNum() map[int]int {
	return tn.LineEntitiesNum
}

// GetText retrieves the text content in UTF-8 runes format.
func (tn *textNER) GetText() []rune {
	return tn.Text
}

// SetEntities saves detected entities into TextNER object.
func (tn *textNER) SetEntities(ents []EntityNER) {
	tn.Entities = ents
}

// GetEntities retrieves detected named entities.
func (tn *textNER) GetEntities() []EntityNER {
	return tn.Entities
}

// ToJSON encodes TextNER into JSON string.
func (tn *textNER) ToJSON(pretty bool) ([]byte, error) {
	enc := gnfmt.GNjson{Pretty: pretty}
	return enc.Encode(tn)
}
