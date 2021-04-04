package token

// Feature represents a property of a Token. Different named entities would
// require a different set of features.
type Feature interface {
	// Analyse examines the token and calculates the Value for the feature.
	Analyse(token *Token)

	// Value returns the value of a feature for the token.
	Value(obj interface{})

	// String is a string representation of a feature.
	String() string
}
