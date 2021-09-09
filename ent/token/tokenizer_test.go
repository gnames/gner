package token_test

import (
	"fmt"
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

	assert.Equal(t, string(res[0].Cleaned()), "h€llö")

	assert.Equal(t, string(res[1].Raw()), "wörl'd")
	assert.Equal(t, res[1].Start(), 6)
	assert.Equal(t, res[1].End(), 12)

	assert.Equal(t, res[1].Cleaned(), "wörl'd")
}

func TestTokenizeBOM(t *testing.T) {
	// BOM at the first position in the text shold be removed before texts is
	// tokenized, but there may be BOM characters in the text from concatenation
	// of texts or from OCR error.
	text := []rune{'*', '\uFEFF', 'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	res := token.Tokenize(text, wrapToken)
	assert.Equal(t, len(res), 3)
	assert.Equal(t, res[1].Line(), 0)
	assert.Equal(t, res[1].Start(), 2)
	assert.Equal(t, res[2].Line(), 0)
	assert.Equal(t, res[2].Start(), 8)
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
	assert.Equal(t, string(res[8].Cleaned()), "1778,")
}

func parseTestdataFile(t *testing.T) []string {
	text, err := os.ReadFile("../../testdata/tokenize.json")
	assert.Nil(t, err)
	return strings.Split(string(text), "\n")
}

func wrapToken(token token.TokenNER) token.TokenNER {
	return token
}

func Example() {
	text := "one\vtwo Poma-  \t\r\n tomus " +
		"dash -\nstandalone " +
		"Tora-\nBora\n\rthree 1778,\n"
	res := token.Tokenize([]rune(text), wrapToken)
	fmt.Println(res[0].Cleaned())
	fmt.Println(res[2].Cleaned())
	// Output:
	// one
	// Pomatomus
}
