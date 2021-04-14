package token_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gnames/gner/ent/token"
	"github.com/stretchr/testify/assert"
)

func TestTokenizeBasic(t *testing.T) {
	text := []rune("hello world")
	res := token.Tokenize(text, wrapToken)
	assert.Equal(t, len(res), 2)

	assert.Equal(t, string(res[0].Raw()), "hello")
	assert.Equal(t, res[0].Line(), 0)
	assert.Equal(t, res[0].Start(), 0)
	assert.True(t, res[0].Properties().IsWord)

	assert.Equal(t, res[1].Start(), 6)
	assert.Equal(t, res[1].End(), 11)
}

func TestTokenizeUTF(t *testing.T) {
	text := []rune("h€llö wörl'd")
	res := token.Tokenize(text, wrapToken)
	assert.Equal(t, len(res), 2)

	assert.Equal(t, string(res[0].Raw()), "h€llö")
	assert.Equal(t, res[0].Line(), 0)
	assert.Equal(t, res[0].Start(), 0)

	assert.Equal(t, string(res[0].Cleaned()), "h�llö")
	assert.False(t, res[0].Properties().IsWord)

	assert.Equal(t, string(res[1].Raw()), "wörl'd")
	assert.Equal(t, res[1].Start(), 6)
	assert.Equal(t, res[1].End(), 12)

	assert.Equal(t, res[1].Cleaned(), "wörl�d")
	assert.True(t, res[1].Properties().HasLetters)
	assert.False(t, res[1].Properties().HasDigits)
	assert.False(t, res[1].Properties().HasDash)
}

func TestTokenizeLines(t *testing.T) {
	text := []rune("h€llö \nwörld")
	res := token.Tokenize(text, wrapToken)
	assert.Equal(t, len(res), 2)
	assert.Equal(t, res[0].Line(), 0)
	assert.Equal(t, res[1].Line(), 1)
}

func TestTokenizeConcatenate(t *testing.T) {
	text := "one\vtwo Poma-  \t\r\n tomus " +
		"dash -\nstandalone " +
		"Tora-\nBora\n\rthree 1778,\n"
	res := token.Tokenize([]rune(text), wrapToken)

	assert.Equal(t, len(res), 9)
	assert.Equal(t, string(res[2].Cleaned()), "Pomatomus")
	assert.Equal(t, string(res[4].Cleaned()), "-")
	assert.Equal(t, string(res[6].Cleaned()), "Tora-Bora")
	assert.True(t, res[6].Properties().IsWord)
	assert.True(t, res[6].Properties().HasDash)
	assert.Equal(t, string(res[8].Cleaned()), "1778")
	assert.True(t, res[8].Properties().IsNumber)
	assert.True(t, res[8].Properties().HasDigits)
}

func TestProperties(t *testing.T) {
	text := "(12oo] die Vögel 1888-1889 Tora-Bora )"
	res := token.Tokenize([]rune(text), wrapToken)

	assert.True(t, res[0].Properties().HasStartParens)
	assert.True(t, res[0].Properties().HasEndSqParens)
	assert.True(t, res[0].Properties().HasDigits)
	assert.True(t, res[0].Properties().HasLetters)
	assert.False(t, res[0].Properties().IsWord)
	assert.True(t, res[2].Properties().IsWord)
	assert.Equal(t, string(res[3].Cleaned()), "1888-1889")
	assert.True(t, res[3].Properties().HasDigits)
	assert.False(t, res[3].Properties().HasLetters)
	assert.True(t, res[3].Properties().HasDash)
	assert.False(t, res[3].Properties().IsWord)
	assert.False(t, res[3].Properties().IsNumber)
	assert.Equal(t, string(res[4].Cleaned()), "Tora-Bora")
	assert.False(t, res[4].Properties().HasDigits)
	assert.True(t, res[4].Properties().HasLetters)
	assert.True(t, res[4].Properties().HasDash)
	assert.True(t, res[4].Properties().IsWord)
	assert.False(t, res[4].Properties().IsNumber)
	assert.Equal(t, string(res[5].Raw()), ")")
	assert.True(t, res[5].Properties().HasEndParens)
	assert.False(t, res[5].Properties().HasStartParens)
	assert.False(t, res[5].Properties().HasLetters)
}

func TestJSON(t *testing.T) {
	data := parseTestdataFile(t)
	text, err := os.ReadFile("../../testdata/tokenize.txt")
	assert.Nil(t, err)
	res := token.Tokenize([]rune(string((text))), wrapToken)
	for i := range res {
		out, err := res[i].ToJSON()
		assert.Nil(t, err)
		assert.Equal(t, string(out), strings.TrimSpace(data[i]))
	}
}

func parseTestdataFile(t *testing.T) []string {
	text, err := os.ReadFile("../../testdata/tokenize.json")
	assert.Nil(t, err)
	return strings.Split(string(text), "\n")
}

func wrapToken(token token.TokenNER) token.TokenNER {
	return token
}
