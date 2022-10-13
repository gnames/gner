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
	assert := assert.New(t)
	text := []rune("hello world")
	res := token.Tokenize(text, wrapToken)
	assert.Equal(2, len(res))
	assert.Equal("hello", string(res[0].Raw()))
	assert.Equal(0, res[0].Line())
	assert.Equal(0, res[0].Start())
	assert.Equal(6, res[1].Start())
	assert.Equal(11, res[1].End())
}

func TestParens(t *testing.T) {
	assert := assert.New(t)
	text := []rune(`d      Family Blaniulidae
        (I) Blaniulus guttulatus (Fabricius, 1798)
        (I) Choneiulus palmatus (Němec, 1895)
      Family Julidae
`)
	res := token.Tokenize(text, wrapToken)
	for i := range res {
		fmt.Printf("TOKEN%d: %#v\n", i, res[i])
	}
	fmt.Println()
	assert.Equal("Family", string(res[0].Cleaned()))
}

func TestTokenizeUTF(t *testing.T) {
	assert := assert.New(t)
	text := []rune("h€llö wörl'd")
	res := token.Tokenize(text, wrapToken)
	assert.Equal(2, len(res))

	assert.Equal("h€llö", string(res[0].Raw()))
	assert.Equal(0, res[0].Line())
	assert.Equal(0, res[0].Start())

	assert.Equal("h€llö", string(res[0].Cleaned()))

	assert.Equal("wörl'd", string(res[1].Raw()))
	assert.Equal(6, res[1].Start())
	assert.Equal(12, res[1].End())

	assert.Equal("wörl'd", res[1].Cleaned())
}

func TestTokenizeBOM(t *testing.T) {
	// BOM at the first position in the text shold be removed before texts is
	// tokenized, but there may be BOM characters in the text from concatenation
	// of texts or from OCR error.
	assert := assert.New(t)
	text := []rune{'*', '\uFEFF', 'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	res := token.Tokenize(text, wrapToken)
	assert.Equal(3, len(res))
	assert.Equal(0, res[1].Line())
	assert.Equal(2, res[1].Start())
	assert.Equal(0, res[2].Line())
	assert.Equal(8, res[2].Start())
}

func TestTokenizeLines(t *testing.T) {
	assert := assert.New(t)
	text := []rune("h€llö \nwörld")
	res := token.Tokenize(text, wrapToken)
	assert.Equal(2, len(res))
	assert.Equal(0, res[0].Line())
	assert.Equal(1, res[1].Line())
}

func TestTokenizeConcatenate(t *testing.T) {
	assert := assert.New(t)
	text := "one\vtwo Poma-  \t\r\n tomus " +
		"dash -\nstandalone " +
		"Tora-\nBora\n\rthree 1778,\n"
	res := token.Tokenize([]rune(text), wrapToken)

	assert.Equal(9, len(res), 9)
	assert.Equal("Pomatomus", string(res[2].Cleaned()))
	assert.Equal("-", string(res[4].Cleaned()))
	assert.Equal("Tora-Bora", string(res[6].Cleaned()))
	assert.Equal("1778,", string(res[8].Cleaned()))
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
	res := token.Tokenize([]rune(text), func(t token.TokenNER) token.TokenNER { return t })
	fmt.Println(res[0].Cleaned())
	fmt.Println(res[2].Cleaned())
	// Output:
	// one
	// Pomatomus
}
