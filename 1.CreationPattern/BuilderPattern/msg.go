package msg

import (
	"sync"
)

/*
创建者模式
在程序设计中，我们会经常遇到一些复杂的对象，其中有很多成员属性，甚至嵌套着多个复杂的对象。这种情况下，创建这个复杂对象就会变得很繁琐。
而对于Go语言来说，最常见的表现就是多层的嵌套实例化：
 obj := &MyObject{
   Field1: &Field1 {
     Param1: &Param1 {
       Val: 0,
    },
     Param2: &Param2 {
       Val: 1,
    },
     ...
  },
   Field2: &Field2 {
     Param3: &Param3 {
       Val: 2,
    },
     ...
  },
   ...
 }
有两个明显的缺点：
1. 对对象使用者不友好，使用者在创建对象时需要知道的细节太多；
2. 代码可读性很差。
针对这种对象成员较多，创建对象逻辑较为繁琐的场景，就适合使用建造者模式来进行优化。
作用:
1. 封装复杂对象的创建过程，使对象使用者不感知复杂的创建逻辑
2. 可以一步步按照顺序对成员进行赋值，或者创建嵌套对象，并最终完成目标对象的创建
3. 对多个对象复用同样的对象创建逻辑
*/

type Message struct {
	Header *Header
	Body   *Body
}

type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}

type Body struct {
	Items []string
}

type builder struct {
	once *sync.Once
	msg  *Message
}

func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg: &Message{
			Header: &Header{},
			Body:   &Body{},
		},
	}
}

func (b *builder) WithSrcAddr(srcAddr string) *builder {
	b.msg.Header.SrcAddr = srcAddr
	return b
}

func (b *builder) WithSrcPort(srcPort uint64) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}

func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}

func (b *builder) WithDestPort(destPort uint64) *builder {
	b.msg.Header.DestPort = destPort
	return b
}

func (b *builder) WithHeaderItem(key, value string) *builder {
	b.once.Do(func() { b.msg.Header.Items = make(map[string]string) })
	b.msg.Header.Items[key] = value
	return b
}

func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

// Build 最后调用来创建Message对象
func (b *builder) Build() *Message {
	return b.msg
}
