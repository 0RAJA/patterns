package factoryMethod

//工厂方法
/*
工厂方法模式跟上一节讨论的建造者模式类似，都是将对象创建的逻辑封装起来，为使用者提供一个简单易用的对象创建接口
1. 提供一个工厂对象，通过调用工厂对象的工厂方法来创建产品对象
*/

type Type uint8

const (
	Start Type = iota
	End
)

type Event interface {
	EventType() Type
	Content() string
}

// StartEvent 开始事件
type StartEvent struct {
	content string
}

func (s *StartEvent) EventType() Type {
	return Start
}

func (s *StartEvent) Content() string {
	return s.content
}

// EndEvent 结束事件
type EndEvent struct {
	content string
}

func (e *EndEvent) EventType() Type {
	return End
}

func (e *EndEvent) Content() string {
	return e.content
}

type Factory struct{}

func (e *Factory) Create(etype Type) Event {
	switch etype {
	case Start:
		return &StartEvent{content: "Start Event"}
	case End:
		return &EndEvent{content: "End Event"}
	}
	return nil
}
