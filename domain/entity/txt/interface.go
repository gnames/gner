package txt

type OutputVolumeNER interface {
	VolumeID() string
	OutputPages() []OutputPageNER
	Formatter
}

type OutputPageNER interface {
	PageID() string
	Formatter
}

type OutputNER interface {
	Formatter
}

type Formatter interface {
	ToJSON(pretty bool) ([]byte, error)
}
