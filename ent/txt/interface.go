package txt

// EntityNER is the most important object for named entity recognition. It
// contains information about an entity found in a text.
type EntityNER interface {
	Formatter
}

// TextNER represents the simplest object for named entities recognition.
// Higher level objects like PageNER and VolumeNER incorporate it.
type TextNER interface {
	// GetText retrievs the raw text represented by UTF-8 encoded runes.
	GetText() []rune

	// SetLinesEntitiesNum stores the number of found named entities per line.
	// The keys are line numbers, the values are numbers of entities for each
	// line.
	SetLinesEntitiesNum(lines map[int]int)

	// GetLinesEntitiesNum retrievs information about number of entities found
	// per line of text.
	GetLinesEntitiesNum() map[int]int

	// SetEntities stores information about all found in the text named entities.
	// Named entities can be scientific names, names of people, geographical
	// places, numbers etc.
	SetEntities(ents []EntityNER)

	// GetEntities retrievs information about all named entities found in the
	// text.
	GetEntities() []EntityNER

	// Formatter interface encodes the data in a format suitable for outputs.
	Formatter
}

// PageNER is an TextNER object that also includes meta-information about the
// page.
type PageNER interface {
	GetID() string
	TextNER
}

// VolumeNER is a book a magazine, a journal etc. It contains 0 or more pages
// and keeps its own meta-information.
type VolumeNER interface {
	GetID() string
	SetPages(pages []PageNER)
	GetPages() []PageNER
	Formatter
}

type Formatter interface {
	ToJSON(pretty bool) ([]byte, error)
}
