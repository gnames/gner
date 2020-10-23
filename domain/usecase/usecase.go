package usecase

import (
	"fmt"
)

type gner struct {
	ner NERecognizer
}

func NewGNER(r NERecognizer) GNER {
	return gner{ner: r}
}

func (g gner) Run() error {
	return nil
}

func (g gner) Version() string {
	ver := fmt.Sprintf("\n\nversion: %s\nbuild: %s\n\n", Version, Build)
	return ver
}
