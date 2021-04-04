package txt

import "github.com/gnames/gnfmt"

// Type page represents a page of a book, magazine, journal, web-page etc.
// It can be stand-alone or be included into a VolumeNER.
type page struct {
	ID string `json:"pageId"`
	TextNER
}

// NNewPageNER is a PageNer factory.
func NewPageNER(id string, text TextNER) PageNER {
	return &page{
		ID:      id,
		TextNER: text,
	}
}

// GetID returns an identifier of a page.
func (p *page) GetID() string {
	return p.ID
}

// ToJSON returns JSON-encoded representation of a PageNER object.
func (p *page) ToJSON(pretty bool) ([]byte, error) {
	enc := gnfmt.GNjson{Pretty: pretty}
	return enc.Encode(p)
}
