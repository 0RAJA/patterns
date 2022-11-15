package main

import (
	"fmt"
)

// WeaponStrategy 抽象武器策略
type WeaponStrategy interface {
	UseWeapon() // 使用武器
}

// 具体策略

type Ak47 struct{}

func (a *Ak47) UseWeapon() {
	fmt.Println("Ak47...")
}

type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("Knife...")
}

// 环境类

type Hero struct {
	strategy WeaponStrategy // 抽象策略
}

func (h *Hero) SetWeaponStrategy(strategy WeaponStrategy) *Hero {
	h.strategy = strategy
	return h
}

// Fight 业务逻辑
func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {
	hero := new(Hero)
	// 策略1
	hero.SetWeaponStrategy(new(Ak47)).Fight()
	// 策略2
	hero.SetWeaponStrategy(new(Knife)).Fight()
}
