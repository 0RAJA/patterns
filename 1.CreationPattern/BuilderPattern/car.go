package msg

import (
	"fmt"
)

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type CarBuilder interface {
	Color(Color) CarBuilder
	Wheels(Wheels) CarBuilder
	TopSpeed(Speed) CarBuilder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type Car struct {
	color    Color
	wheels   Wheels
	topSpeed Speed
}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) Drive() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%v", p)
		}
	}()
	fmt.Println(c.color, c.wheels, c.topSpeed, "is running")
	return
}

func (c *Car) Stop() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%v", p)
		}
	}()
	fmt.Println(c.color, c.wheels, c.topSpeed, "is done")
	return
}

func (c *Car) Color(color Color) CarBuilder {
	c.color = color
	return c
}

func (c *Car) Wheels(wheels Wheels) CarBuilder {
	c.wheels = wheels
	return c
}

func (c *Car) TopSpeed(speed Speed) CarBuilder {
	c.topSpeed = speed
	return c
}

func (c *Car) Build() Interface {
	return c
}
