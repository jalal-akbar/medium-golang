package oop

import (
	"fmt"
	"testing"
)

type instrument interface {
	play()
}

func playInstrument[I instrument](i I) {
	i.play()
}

type guitar struct {
	Sounds string
}

func (g guitar) play() {
	fmt.Println("Guitar Sounds")
}

func TestGenerics(t *testing.T) {
	g := guitar{"Guitar Sound"}
	playInstrument(g)
}
