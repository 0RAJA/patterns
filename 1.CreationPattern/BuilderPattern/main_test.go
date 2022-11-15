package msg_test

import (
	"testing"

	msg "patterns/1.CreationPattern/BuilderPattern"
)

// 直接创建
func sample1() {
	message := msg.Message{
		Header: &msg.Header{
			SrcAddr:  "",
			SrcPort:  0,
			DestAddr: "",
			DestPort: 0,
			Items:    nil,
		},
		Body: &msg.Body{Items: make([]string, 0)},
	}
	message.Header.Items["contents"] = "application/json"
	message.Body.Items = append(message.Body.Items, "record1")
	message.Body.Items = append(message.Body.Items, "record2")
}

func TestSample2(t *testing.T) {
	message := msg.Builder().WithSrcAddr("127.0.0.1").WithSrcPort(8080).WithDestAddr("127.0.0.2").WithDestPort(8080).WithHeaderItem("contents", "application/json").WithBodyItem("record1").WithBodyItem("record2").Build()
	if message.Header.SrcAddr != "127.0.0.1" {
		t.Errorf("expect src address 192.168.0.1, but actual %s.", message.Header.SrcAddr)
	}
	if message.Body.Items[0] != "record1" {
		t.Errorf("expect body item0 record1, but actual %s.", message.Body.Items[0])
	}
}

func TestSample3(t *testing.T) {
	assembly := msg.NewCar().Color(msg.RedColor)

	familyCar := assembly.Wheels(msg.SportsWheels).TopSpeed(50 * msg.MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(msg.SteelWheels).TopSpeed(150 * msg.MPH).Build()
	sportsCar.Drive()

	familyCar.Stop()
	sportsCar.Stop()
}
