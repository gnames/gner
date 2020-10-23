package usecase

type NERecognizer interface {
	Init() error
	Find(text []rune, output interface{}) error
	Verify(input interface{}, output interface{}) error
	Version()
}

type GNER interface {
	Run() error
	Version() string
}
