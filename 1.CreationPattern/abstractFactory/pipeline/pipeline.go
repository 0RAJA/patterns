package pipeline

import "patterns/1.CreationPattern/abstractFactory/plugin"

// Pipeline 消息管道的定义
type Pipeline struct {
	input  plugin.Input
	filter plugin.Filter
	output plugin.Output
}

// Exec 消息的处理流程 input->filter->output
func (p *Pipeline) Exec() {
	msg := p.input.Receive()
	msg = p.filter.Process(msg)
	p.output.Send(msg)
}
