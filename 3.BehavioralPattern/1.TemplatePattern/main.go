package main

import (
	"fmt"
)

// Make 模板方法
type Make interface {
	BoilWater()
	Brew()
	PourInCup()
	AddThings()
}

// Template 模板
type Template struct {
	b Make
}

func NewTemplate(make Make) *Template {
	return &Template{b: make}
}

// Make 模板方法
func (t *Template) Make() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	t.b.AddThings()
}

// 具体类

// MakeCoffee 制作coffee
type MakeCoffee struct{}

func (m *MakeCoffee) BoilWater() {
	fmt.Println("水-100")
}

func (m *MakeCoffee) Brew() {
	fmt.Println("水冲咖啡")
}

func (m *MakeCoffee) PourInCup() {
	fmt.Println("咖啡到杯子")
}

func (m *MakeCoffee) AddThings() {
	fmt.Println("添加一些东西")
}

// MakeTee 制作茶
type MakeTee struct{}

func (m *MakeTee) BoilWater() {
	fmt.Println("water-tee")
}

func (m *MakeTee) Brew() {
	fmt.Println("brew-tee")
}

func (m *MakeTee) PourInCup() {
	fmt.Println("pourInCup-tee")
}

func (m *MakeTee) AddThings() {
	fmt.Println("AddThing-tee")
}

func main() {
	// 1. 制作Coffee
	makeCoffee := NewTemplate(new(MakeCoffee))
	makeCoffee.Make()
	// 2. 制作tee
	makeTee := NewTemplate(new(MakeTee))
	makeTee.Make()
}
