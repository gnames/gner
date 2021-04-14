package token

import (
	"strings"
	"unicode"

	"github.com/gnames/gnfmt"
)

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

	// properties is a fixed set of general properties that we determine during
	// the text traversal.
	properties *Properties
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

	// IsCapitalized is true if the furst letter of a token is capitalized.
	// The first letter does not have to be the first character.
	IsCapitalized bool

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

func (t *tokenNER) Properties() *Properties {
	return t.properties
}

func (t *tokenNER) SetProperties(p *Properties) {
	t.properties = p
}

// CalculateProperties takes raw and cleaned values of a token and computes
// properties of these values, saving them into Properties object.
func CalculateProperties(raw, cleaned []rune, p *Properties) {
	p.HasStartParens = raw[0] == rune('(')
	p.HasEndParens = raw[len(raw)-1] == rune(')')
	p.HasStartSqParens = raw[0] == rune('[')
	p.HasEndSqParens = raw[len(raw)-1] == rune(']')
	p.HasEndDot = raw[len(raw)-1] == rune('.')
	p.HasEndComma = raw[len(raw)-1] == rune(',')
	for _, v := range cleaned {
		if v == rune('-') {
			p.HasDash = true
		}

		if !p.HasLetters && unicode.IsLetter(v) {
			p.HasLetters = true
			continue
		}

		if !p.HasDigits && unicode.IsDigit(v) {
			p.HasDigits = true
			continue
		}

		if !p.HasSpecialChars && v == rune('�') {
			p.HasSpecialChars = true
			continue
		}
	}

	if p.HasDigits && !p.HasLetters && !p.HasSpecialChars && !p.HasDash {
		p.IsNumber = true
	}

	if p.HasLetters && !p.HasDigits && !p.HasSpecialChars {
		p.IsWord = true
	}
}

func (t *tokenNER) ProcessRaw() {
	t.normalizeRaw()
	t.properties = &Properties{}
	CalculateProperties(t.raw, []rune(t.cleaned), t.properties)
}

func (t *tokenNER) normalizeRaw() {
	var runes []rune
	t.runeSet = make(map[rune]struct{})
	firstLetter := true
	for i, v := range t.raw {
		t.runeSet[v] = struct{}{}
		hasDash := v == rune('-')
		if unicode.IsLetter(v) || unicode.IsNumber(v) || hasDash {
			if firstLetter {
				t.cleanedStart = i
				firstLetter = false
			}
			t.cleanedEnd = i
			runes = append(runes, v)
		} else {
			t.runeSet['�'] = struct{}{}
			runes = append(runes, rune('�'))
		}
	}
	res := string(runes)
	t.cleaned = strings.Trim(res, "�")
}

// ToJSON serializes token to JSON string
func (t *tokenNER) ToJSON() ([]byte, error) {
	enc := gnfmt.GNjson{}
	tj := TokenJSON{
		Line:    t.line,
		Raw:     string(t.raw),
		Cleaned: string(t.cleaned),
		Start:   t.start,
		End:     t.end,
	}
	return enc.Encode(tj)
}
