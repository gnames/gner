package txt

import "github.com/gnames/gnlib/encode"

type page struct {
	ID       string      `json:"pageId"`
	Text     []rune      `json:"-"`
	Lines    map[int]int `json:"lines"`
	Entities []EntityNER `json:"entities"`
}

func NewPageNER(id string, text []rune) PageNER {
	return &page{
		ID:   id,
		Text: text,
	}
}

func (p *page) GetID() string {
	return p.ID
}

func (p *page) GetText() []rune {
	return p.Text
}

func (p *page) SetLines(lines map[int]int) {
	p.Lines = lines
}

func (p *page) GetLines() map[int]int {
	return p.Lines
}

func (p *page) SetEntities(ents []EntityNER) {
	p.Entities = ents
}

func (p *page) GetEntities() []EntityNER {
	return p.Entities
}

func (p *page) ToJSON(pretty bool) ([]byte, error) {
	enc := encode.GNjson{Pretty: pretty}
	return enc.Encode(p)
}
