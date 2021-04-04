package token_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gnames/gner/ent/token"
	"github.com/stretchr/testify/assert"
)

func TestTokenizeBasic(t *testing.T) {
	text := []rune("hello world")
	res := token.Tokenize(text)
	assert.Equal(t, len(res), 2)
	assert.Equal(t, string(res[0].Cleaned), "hello")
	assert.Equal(t, res[0].Line, 0)
	assert.Equal(t, res[0].Start, 0)
	assert.True(t, res[0].IsWord)
	assert.Equal(t, res[1].Start, 6)
	assert.Equal(t, res[1].End, 11)
}

func TestTokenizeUTF(t *testing.T) {
	text := []rune("h€llö wörl'd")
	res := token.Tokenize(text)
	assert.Equal(t, len(res), 2)
	assert.Equal(t, string(res[0].Raw), "h€llö")
	assert.Equal(t, string(res[0].Cleaned), "h�llö")
	assert.Equal(t, res[0].Line, 0)
	assert.Equal(t, res[0].Start, 0)
	assert.False(t, res[0].IsWord)
	assert.Equal(t, res[1].Start, 6)
	assert.Equal(t, res[1].End, 12)
	assert.Equal(t, res[1].Cleaned, "wörl�d")
	assert.True(t, res[1].HasLetters)
	assert.False(t, res[1].HasDigits)
	assert.False(t, res[1].HasDash)
}

func TestTokenizeLines(t *testing.T) {
	text := []rune("h€llö \nwörld")
	res := token.Tokenize(text)
	assert.Equal(t, len(res), 2)
	assert.Equal(t, res[0].Line, 0)
	assert.Equal(t, res[1].Line, 1)
}

func TestTokenizeConcatenate(t *testing.T) {
	text := "one\vtwo Poma-  \t\r\n tomus " +
		"dash -\nstandalone " +
		"Tora-\nBora\n\rthree 1778,\n"
	res := token.Tokenize([]rune(text))
	assert.Equal(t, len(res), 9)
	assert.Equal(t, string(res[2].Cleaned), "Pomatomus")
	assert.Equal(t, string(res[4].Cleaned), "-")
	assert.Equal(t, string(res[6].Cleaned), "Tora-Bora")
	assert.True(t, res[6].IsWord)
	assert.True(t, res[6].HasDash)
	assert.Equal(t, string(res[8].Cleaned), "1778")
	assert.True(t, res[8].IsNumber)
	assert.True(t, res[8].HasDigits)
}

func TestProperties(t *testing.T) {
	text := "(12oo] die Vögel 1888-1889 Tora-Bora )"
	res := token.Tokenize([]rune(text))
	assert.True(t, res[0].HasStartParens)
	assert.True(t, res[0].HasEndSqParens)
	assert.True(t, res[0].HasDigits)
	assert.True(t, res[0].HasLetters)
	assert.False(t, res[0].IsWord)
	assert.True(t, res[2].IsWord)
	assert.Equal(t, string(res[3].Cleaned), "1888-1889")
	assert.True(t, res[3].HasDigits)
	assert.False(t, res[3].HasLetters)
	assert.True(t, res[3].HasDash)
	assert.False(t, res[3].IsWord)
	assert.False(t, res[3].IsNumber)
	assert.Equal(t, string(res[4].Cleaned), "Tora-Bora")
	assert.False(t, res[4].HasDigits)
	assert.True(t, res[4].HasLetters)
	assert.True(t, res[4].HasDash)
	assert.True(t, res[4].IsWord)
	assert.False(t, res[4].IsNumber)
	assert.Equal(t, string(res[5].Raw), ")")
	assert.True(t, res[5].HasEndParens)
	assert.False(t, res[5].HasStartParens)
	assert.False(t, res[5].HasLetters)
}

func TestJSON(t *testing.T) {
	data := parseTestdataFile(t)
	text, err := ioutil.ReadFile("../../testdata/tokenize.txt")
	assert.Nil(t, err)
	res := token.Tokenize([]rune(string((text))))
	for i, token := range res {
		out, err := token.ToJson()
		assert.Nil(t, err)
		assert.Equal(t, string(out), strings.TrimSpace(data[i]))
	}
}

func parseTestdataFile(t *testing.T) []string {
	text, err := ioutil.ReadFile("../../testdata/tokenize.json")
	assert.Nil(t, err)
	return strings.Split(string(text), "\n")
}