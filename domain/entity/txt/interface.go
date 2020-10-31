package txt

type VolumeNER interface {
	GetID() string
	SetPages(pages []PageNER)
	GetPages() []PageNER
	Formatter
}

type PageNER interface {
	GetID() string
	TextNER
}

type TextNER interface {
	GetText() []rune

	SetLines(lines map[int]int)
	GetLines() map[int]int

	SetEntities(ents []EntityNER)
	GetEntities() []EntityNER

	Formatter
}

type EntityNER interface {
	Formatter
}

type Formatter interface {
	ToJSON(pretty bool) ([]byte, error)
}
