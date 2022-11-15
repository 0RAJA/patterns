package plugin

import (
	"fmt"
	"reflect"
	"strings"
)

/*
抽象工厂模式
存在问题:
	在工厂方法模式中，我们通过一个工厂对象来创建一个产品族，具体创建哪个产品，则通过swtich-case的方式去判断。这也意味着该产品组上，每新增一类产品对象，都必须修改原来工厂对象的代码；而且随着产品的不断增多，工厂对象的职责也越来越重，违反了单一职责原则。
解决方法:
	抽象工厂模式通过给工厂类新增一个抽象层解决了该问题
*/
/*
考虑需要如下一个插件架构风格的消息处理系统，
pipeline是消息处理的管道，其中包含了input、filter和output三个插件。
我们需要实现根据配置来创建pipeline ，加载插件过程的实现非常适合使用工厂模式，其中input、filter和output三类插件的创建使用抽象工厂模式，而pipeline的创建则使用工厂方法模式。
*/

type Plugin interface{}

// Input 输入插件,接收消息
type Input interface {
	Plugin
	Receive() string
}

// Filter 过滤插件,处理消息
type Filter interface {
	Plugin
	Process(msg string) string
}

// Output 输出插件,发送消息
type Output interface {
	Plugin
	Send(msg string)
}

//插件名与类型的映射,通过反射创建对象
var (
	inputNames  = make(map[string]reflect.Type)
	filterNames = make(map[string]reflect.Type)
	outputNames = make(map[string]reflect.Type)
)

func init() {
	inputNames["hello"] = reflect.TypeOf(InputHello{})
}

func init() {
	filterNames["upper"] = reflect.TypeOf(FilterUpper{})
}

func init() {
	outputNames["console"] = reflect.TypeOf(OutputConsole{})
}

// InputHello 接收插件
type InputHello struct{}

func (i *InputHello) Receive() string {
	return "hello"
}

// FilterUpper 处理插件
type FilterUpper struct{}

func (f *FilterUpper) Process(msg string) string {
	return strings.ToUpper(msg)
}

// OutputConsole 输出插件
type OutputConsole struct{}

func (o *OutputConsole) Send(msg string) {
	fmt.Println(msg)
}
