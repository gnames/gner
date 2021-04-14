package token

// TokenNER represents a word separated by spaces in a text. Words split by new
// lines are concatenated.
type TokenNER interface {
	// Raw is a verbatim presentation of a token as it appears in a text.
	Raw() []rune

	// Start is the index of the first rune of a token. The first rune
	// does not have to be alpha-numeric.
	Start() int

	// End is the index of the last rune of a token. The last rune does not
	// have to be alpha-numeric.
	End() int

	// Line line number in the text
	Line() int

	// SetLine sets the line number
	SetLine(int)

	// Cleaned is a presentation of a token after normalization.
	Cleaned() string

	// SetCleaned substitues existing cleaned text with a new one.
	SetCleaned(string)

	// Properties is a fixed set of general properties that we determine during
	// the text traversal.
	Properties() *Properties

	// SetProperties substitutes existing properties with new ones.
	SetProperties(*Properties)

	// ProcessRaw computes a clean version of a name as well as properties
	// of the token.
	ProcessRaw()

	// ToJSON converts TokenNER object into JSON represenation
	ToJSON() ([]byte, error)
}
