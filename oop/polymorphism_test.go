package oop

import (
	"fmt"
	"testing"
)

type Instrument interface {
	Play()
}

func PlayInstrument(i Instrument) {
	i.Play()
}

type Guitar struct {
	Type string
}

func (g Guitar) Play() {
	fmt.Println("Guitar Sounds")
}

func TestPolymorphism(t *testing.T) {
	g := Guitar{
		Type: "Classical",
	}
	PlayInstrument(g)
}
