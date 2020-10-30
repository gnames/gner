package ner

import "github.com/gnames/gner/domain/entity/txt"

type NERecognizer interface {
	Find(text []rune) []txt.OutputNER
	FindInVolume(volume txt.Volume) txt.OutputVolumeNER
}
