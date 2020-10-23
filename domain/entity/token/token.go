package token

import (
	"strings"
	"unicode"

	"github.com/gnames/gnlib/encode"
)

// Token represents a word separated by spaces in a text. Words split by new
// lines are concatenated.
type Token struct {
	// Line line number in the text
	Line int

	// Raw is a verbatim presentation of a token as it appears in a text.
	Raw []rune

	// Start is the index of the first rune of a token. The first rune
	// does not have to be alpha-numeric.
	Start int

	// End is the index of the last rune of a token. The last rune does not
	// have to be alpha-numeric.
	End int

	// runeSet provides runes that exist in the token
	runeSet map[rune]struct{}

	// Cleaned is a presentation of a token after normalization.
	Cleaned string

	// cleanedStart is the first rune of cleaned token
	cleanedStart int

	// cleanedEnd is the last rune of clenaed token
	cleanedEnd int

	// Features is the map of features as values with their string
	// representations as keys.
	Features map[string]Feature
}

//TokenJSON provides a presentation view for a Token.
type TokenJSON struct {
	Line    int    `json:"lineNumber"`
	Raw     string `json:"raw"`
	Cleaned string `json:"cleaned"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
}

// NewToken constructs a new Token object.
func NewToken(text []rune, start int, end int, feat ...Feature) Token {
	t := Token{
		Raw:     text[start:end],
		Start:   start,
		End:     end,
		runeSet: make(map[rune]struct{}),
	}
	t.clean()

	for _, feature := range feat {
		feature.Analyse(&t)
		t.Features[feature.String()] = feature
	}
	return t
}

// Clean converts a verbatim (Raw) string of a token into normalized cleaned up
// version.
func (t *Token) clean() {
	var res []rune
	firstLetter := true
	for i, v := range t.Raw {
		t.runeSet[v] = struct{}{}
		hasDash := v == rune('-')
		if unicode.IsLetter(v) || unicode.IsNumber(v) || hasDash {
			if firstLetter {
				t.cleanedStart = i
				firstLetter = false
			}
			t.cleanedEnd = i
			res = append(res, v)
		} else {
			t.runeSet['�'] = struct{}{}
			res = append(res, rune('�'))
		}
	}
	t.Cleaned = string(res)
	t.Cleaned = strings.Trim(t.Cleaned, "�")
}

// ToJSON serializes token to JSON string
func (t *Token) ToJson() ([]byte, error) {
	enc := encode.GNjson{}
	tj := TokenJSON{
		Line:    t.Line,
		Raw:     string(t.Raw),
		Cleaned: string(t.Cleaned),
		Start:   t.Start,
		End:     t.End,
	}
	return enc.Encode(tj)
}
