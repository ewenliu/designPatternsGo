package creational

import (
	"fmt"
	"strings"
)

type Vehicle interface {
	Work()
}

type Car struct {
}

func (c *Car) Work() {
	fmt.Println("Car has four wheels")
}

type Bicycle struct {
}

func (b *Bicycle) Work() {
	fmt.Println("Bicycle has two wheels")
}

type AirPlane struct {
}

func (a *AirPlane) Work() {
	fmt.Println("Airplane doesn't has wheels but it can fly")
}

func ChooseBicycle(vehicle string) Vehicle {
	switch strings.ToUpper(vehicle){
	case "CAR":
		return new(Car)
	case "BICYCLE":
		return new(Bicycle)
	case "AIRPLANE":
		return new(AirPlane)
	default:
		panic("No such vehicle!")
	}
}
