package main

import (
	"fmt"
)

// Command 抽象命令
type Command interface {
	Treat()
}

// Doctor 核心执行者
type Doctor struct{}

func (d *Doctor) TreatEye() {
	fmt.Println("treat eye...")
}

func (d *Doctor) TreatNose() {
	fmt.Println("treat nose...")
}

// CommandTreatEye 具体命令
type CommandTreatEye struct {
	doctor *Doctor
}

func (c *CommandTreatEye) Treat() {
	c.doctor.TreatEye()
}

// CommandTreatNose 具体命令
type CommandTreatNose struct {
	doctor *Doctor
}

func (c *CommandTreatNose) Treat() {
	c.doctor.TreatNose()
}

// Nurser 集中管理并执行抽象命令
type Nurser struct {
	cmdList []Command // 命令s
}

// Append 收集命令
func (n *Nurser) Append(cmds ...Command) {
	n.cmdList = append(n.cmdList, cmds...)
}

// Notify 通知执行命令
func (n *Nurser) Notify() {
	for _, cmd := range n.cmdList {
		cmd.Treat()
	}
}

func main() {
	// 执行者
	doctor := new(Doctor)
	// 具体指令
	cmdEye := &CommandTreatEye{doctor: doctor}
	cmdNose := &CommandTreatNose{doctor: doctor}
	// 中间者集中管理和通知执行指令
	nurse := new(Nurser)
	nurse.Append(cmdEye, cmdNose)
	nurse.Notify()
}
