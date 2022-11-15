package main

import (
	"fmt"
)

// 抽象接口

// Attack 主服务
type Attack interface {
	Fight()
}

// PowerOff 适配者--第三方服务
type PowerOff interface {
	Shutdown()
}

// 具体类

type Hero struct {
	Name string
	Arms Attack
}

func (h *Hero) Do() {
	fmt.Println("使用武器...")
	h.Arms.Fight()
}

func NewHero(name string, arms Attack) *Hero {
	return &Hero{name, arms}
}

// Sword 实现武器
type Sword struct{}

func (s *Sword) Fight() {
	fmt.Println("Sword...")
}

// CommonPower 适配者
type CommonPower struct{}

func (p *CommonPower) Shutdown() {
	fmt.Println("shutdown...")
}

// AdaptorForPowerOff 适配第三方服务到主服务
type AdaptorForPowerOff struct {
	p PowerOff
}

func NewAdaptorForPowerOff(p PowerOff) *AdaptorForPowerOff {
	return &AdaptorForPowerOff{p: p}
}

func (f *AdaptorForPowerOff) Fight() {
	fmt.Println("适配器...")
	f.p.Shutdown()
}

func main() {
	hero1 := NewHero("one", &Sword{})
	hero1.Do()
	hero2 := NewHero("two", NewAdaptorForPowerOff(&CommonPower{}))
	hero2.Do()
}
