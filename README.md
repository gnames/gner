# gner

[![DOI](https://zenodo.org/badge/306673142.svg)](https://doi.org/10.5281/zenodo.14262551)

Global Named Entity Recognition

This project provides libraries that a shared between different named entity
recongition (NER) projects. In the future it might also have a command line
application that combines functionality of several NER projects.

## Usage

```go
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
```
