package oop

import (
	"fmt"
	"testing"
)

type Vehicle struct {
	Seats int
	Color string
}

type Car struct {
	Vehicle
}

type MotorCycle struct {
	Base Vehicle
}

func TestInheritance(t *testing.T) {
	car := &Car{
		Vehicle{
			Seats: 4,
			Color: "Blue",
		},
	}

	fmt.Println(car.Seats)
	fmt.Println(car.Color)

	motorCycle := &MotorCycle{
		Base: Vehicle{
			Seats: 3,
			Color: "Green",
		},
	}
	fmt.Println(motorCycle.Base.Seats)
	fmt.Println(motorCycle.Base.Color)
}
