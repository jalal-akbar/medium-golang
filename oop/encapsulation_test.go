package oop

import (
	"fmt"
	"testing"
)

type customer struct {
	id   int
	name string
}

func (c *customer) GetID() int {
	return c.id
}

func (c *customer) GetName() string {
	return c.name
}

func TestEncapsulation(t *testing.T) {
	c := &customer{
		id:   1,
		name: "jalal",
	}
	id := c.GetID()
	name := c.GetName()
	fmt.Println(id)
	fmt.Println(name)
}
