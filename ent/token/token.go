package token

import (
	"strings"
	"unicode"

	"github.com/gnames/gnfmt"
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

	// Properties is a fixed set of general properties that we determine during
	// the text traversal.
	Properties

	// Features is the map of features as values with their string
	// representations as keys.
	Features map[string]Feature
}

// Properties is a fixed set of general properties determined during the
// the text traversal.
type Properties struct {
	// HasStartParens token starts with '('.
	HasStartParens bool

	// HasEndParens token end with '('.
	HasEndParens bool

	// HasStartSqParens token starts with '['.
	HasStartSqParens bool

	// HasEndSqParens token ends with ']'.
	HasEndSqParens bool

	// HasEndDot token ends with '.'
	HasEndDot bool

	// HasEndComma token ends with ','
	HasEndComma bool

	// HasDigits token includes at least one '0-9'.
	HasDigits bool

	// HasLetters token includes at least one character for which
	// unicode.IsLetter(ch) is true.
	HasLetters bool

	// HasDash token includes '-'
	HasDash bool

	// HasSpecialChars internal part of a token includes non-letters, non-digits.
	HasSpecialChars bool

	// IsNumber internal part of a token has only numbers.
	IsNumber bool

	// IsWord internal part of a token includes only letters.
	IsWord bool
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
func NewToken(raw []rune, start int, end int, feat ...Feature) Token {
	t := Token{
		Raw:     raw,
		Start:   start,
		End:     end,
		runeSet: make(map[rune]struct{}),
	}
	t.clean()
	t.properties()

	for _, feature := range feat {
		feature.Analyse(&t)
		t.Features[feature.String()] = feature
	}
	return t
}

// properties examines token and determines its properties.
func (t *Token) properties() {
	t.HasStartParens = t.Raw[0] == rune('(')
	t.HasEndParens = t.Raw[len(t.Raw)-1] == rune(')')
	t.HasStartSqParens = t.Raw[0] == rune('[')
	t.HasEndSqParens = t.Raw[len(t.Raw)-1] == rune(']')
	t.HasEndDot = t.Raw[len(t.Raw)-1] == rune('.')
	t.HasEndComma = t.Raw[len(t.Raw)-1] == rune(',')
	for _, v := range t.Cleaned {
		if v == rune('-') {
			t.HasDash = true
		}

		if !t.HasLetters && unicode.IsLetter(v) {
			t.HasLetters = true
			continue
		}

		if !t.HasDigits && unicode.IsDigit(v) {
			t.HasDigits = true
			continue
		}

		if !t.HasSpecialChars && v == rune('�') {
			t.HasSpecialChars = true
			continue
		}
	}

	if t.HasDigits && !t.HasLetters && !t.HasSpecialChars && !t.HasDash {
		t.IsNumber = true
		return
	}

	if t.HasLetters && !t.HasDigits && !t.HasSpecialChars {
		t.IsWord = true
	}
}

// clean converts a verbatim (Raw) string of a token into normalized cleaned up
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
	enc := gnfmt.GNjson{}
	tj := TokenJSON{
		Line:    t.Line,
		Raw:     string(t.Raw),
		Cleaned: string(t.Cleaned),
		Start:   t.Start,
		End:     t.End,
	}
	return enc.Encode(tj)
}
