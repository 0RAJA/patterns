package main

import (
	"fmt"
)

// 外观模式

type Light struct{}

func (l *Light) On() {
	fmt.Println("light on...")
}

func (l *Light) Off() {
	fmt.Println("light off...")
}

type Xbox struct{}

func (x *Xbox) On() {
	fmt.Println("xbox on...")
}

func (x *Xbox) Off() {
	fmt.Println("xbox off...")
}

// 外观模式,组合上述功能，提供组合选项

type Home struct {
	xbox  *Xbox
	light *Light
}

func (h *Home) Game() {
	h.light.Off()
	h.xbox.On()
}

func (h *Home) Work() {
	h.xbox.Off()
	h.light.On()
}

func main() {
	home := new(Home)
	home.Game()
	home.Work()
}
