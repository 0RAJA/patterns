package main

import (
	"fmt"
)

// Phone ---- 抽象层 ----
type Phone interface {
	Show() string // 构建基本功能
}

// Decorator 抽象的装饰器,装饰器的基础类,此类本应该是interface，但golang的语法不可以有成员变量
type Decorator struct {
	phone Phone // 组合进来抽象的Phone
}

// ---- 实现层 ----
// 具体的构件

type Huawei struct{}

func (h *Huawei) Show() string {
	return "I am Huawei..."
}

type Xiaomi struct{}

func (x *Xiaomi) Show() string {
	return "I am Xiaomi..."
}

// 具体的装饰器

type MoDecorator struct {
	d Decorator
}

func (m *MoDecorator) Show() string {
	return "--mo--" + m.d.phone.Show()
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

type ShellDecorator struct {
	d Decorator
}

func (s *ShellDecorator) Show() string {
	return "--shell--" + s.d.phone.Show()
}

func NewShellDecorator(phone Phone) Phone {
	return &ShellDecorator{Decorator{phone: phone}}
}

// ---- 业务逻辑层 ----
func main() {
	// 1
	var huawei Phone = new(Huawei)
	moHuawei := NewMoDecorator(huawei)
	shellMoHuawei := NewShellDecorator(moHuawei)
	fmt.Println(shellMoHuawei.Show())
	// 2.
	var xiaomi Phone = new(Xiaomi)
	shellXiaomi := NewShellDecorator(xiaomi)
	moShellXiaomi := NewMoDecorator(shellXiaomi)
	fmt.Println(moShellXiaomi.Show())
}
