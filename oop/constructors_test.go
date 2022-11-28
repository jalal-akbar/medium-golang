package oop

import (
	"fmt"
	"testing"
)

type Customer struct {
	Name string
	Age  int
}

func (c *Customer) NewCustomer(name string, age int) {
	c.Name = name
	c.Age = age
}

func TestConstructors(t *testing.T) {
	customer := &Customer{}
	customer.NewCustomer("Jalal", 26)

	fmt.Printf("Name: %s\nAge: %d\n", customer.Name, customer.Age)
}
