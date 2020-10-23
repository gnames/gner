package token

type Feature interface {
	Analyse(token *Token)
	Value(obj interface{})
	String() string
}
