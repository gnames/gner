package txt

import "github.com/gnames/gnlib/encode"

type volume struct {
	ID    string    `json:"volumeId"`
	Pages []PageNER `json:"pages"`
}

func NewVolumeNER(id string) VolumeNER {
	return &volume{ID: id}
}

func (v *volume) GetID() string {
	return v.ID
}

func (v *volume) SetPages(pages []PageNER) {
	v.Pages = pages
}

func (v *volume) GetPages() []PageNER {
	return v.Pages
}

func (v *volume) ToJSON(pretty bool) ([]byte, error) {
	enc := encode.GNjson{Pretty: pretty}
	return enc.Encode(v)
}
