package token

// tokenNER represents a word separated by spaces in a text. Words split by new
// lines are concatenated.
type tokenNER struct {
	// line line number in the text
	line int

	// raw is a verbatim presentation of a token as it appears in a text.
	raw []rune

	// start is the index of the first rune of a token. The first rune
	// does not have to be alpha-numeric.
	start int

	// end is the index of the last rune of a token. The last rune does not
	// have to be alpha-numeric.
	end int

	// runeSet provides runes that exist in the token
	runeSet map[rune]struct{}

	// cleaned is a presentation of a token after normalization.
	cleaned string

	// cleanedStart is the first rune of cleaned token
	cleanedStart int

	// cleanedEnd is the last rune of clenaed token
	cleanedEnd int
}

func (t *tokenNER) Line() int {
	return t.line
}

func (t *tokenNER) SetLine(i int) {
	t.line = i
}

func (t *tokenNER) Raw() []rune {
	return t.raw
}

func (t *tokenNER) Start() int {
	return t.start
}

func (t *tokenNER) End() int {
	return t.end
}

func (t *tokenNER) Cleaned() string {
	return t.cleaned
}

func (t *tokenNER) SetCleaned(s string) {
	t.cleaned = s
}

// ProcessToken provides a hook to implement token-specific logic.
func (t *tokenNER) ProcessToken() {
	t.cleaned = string(t.raw)
}
