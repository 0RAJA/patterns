package main

import (
	"fmt"
)

// 抽象层

// Observer 观察者
type Observer interface {
	DoThing()
}

// Notifier 通知者(被观察者)
type Notifier interface {
	Attach(obs ...Observer)
	Detach(obs Observer)
	Notify()
}

// 实现层

type StuZhangSan struct {
	BadThing string
}

func (s *StuZhangSan) DoThing() {
	fmt.Println("i am ZhangSan,do:", s.BadThing)
}

type StuWangWu struct {
	BadThing string
}

func (s *StuWangWu) DoThing() {
	fmt.Println("i am Wangwu,do:", s.BadThing)
}

type ClassMonitor struct {
	observers map[Observer]struct{}
}

func NewClassMonitor() *ClassMonitor {
	return &ClassMonitor{observers: map[Observer]struct{}{}}
}

func (c *ClassMonitor) Attach(observers ...Observer) {
	for _, obs := range observers {
		c.observers[obs] = struct{}{}
	}
}

func (c *ClassMonitor) Detach(obs Observer) {
	delete(c.observers, obs)
}

func (c *ClassMonitor) Notify() {
	for obs := range c.observers {
		obs.DoThing() // 多态现象
	}
}

// 逻辑层
func main() {
	// 观察者
	var s1 Observer = &StuWangWu{BadThing: "eat"}
	var s2 Observer = &StuZhangSan{BadThing: "sleep"}
	// 通知者
	var monitor Notifier = NewClassMonitor()
	monitor.Attach(s1, s2)
	monitor.Notify()
}
