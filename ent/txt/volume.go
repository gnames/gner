package txt

import "github.com/gnames/gnfmt"

// volume represents a book, a magazine, a journal etc.
type volume struct {
	ID    string    `json:"volumeId"`
	Pages []PageNER `json:"pages"`
}

// NeNewVolumeNER is a factory for VolumeNER object.
func NewVolumeNER(id string) VolumeNER {
	return &volume{ID: id}
}

// GetID returns an identifier of a volume.
func (v *volume) GetID() string {
	return v.ID
}

// SetPages saves PageNER objects created from the volume's content.
func (v *volume) SetPages(pages []PageNER) {
	v.Pages = pages
}

// GetPages returns PageNER objects generated from the volume's content.
func (v *volume) GetPages() []PageNER {
	return v.Pages
}

// ToJSON encodes volume information into JSON.
func (v *volume) ToJSON(pretty bool) ([]byte, error) {
	enc := gnfmt.GNjson{Pretty: pretty}
	return enc.Encode(v)
}
