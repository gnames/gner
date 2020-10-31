package ner

import "github.com/gnames/gner/domain/entity/txt"

type NERecognizer interface {
	Find(tn txt.TextNER)
	FindInVolume(vn txt.VolumeNER)
}
