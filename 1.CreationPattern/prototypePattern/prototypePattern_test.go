package msg_test

import (
	"testing"

	msg "patterns/1.CreationPattern/prototypePattern"
)

func TestBuilder(t *testing.T) {
	message := msg.Builder().
		WithSrcAddr("192.168.0.1").
		WithSrcPort(1234).
		WithDestAddr("192.168.0.2").
		WithDestPort(8080).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Build()
	// 克隆新的
	newMessage := message.Clone().(*msg.Message)
	if newMessage.Header.SrcAddr != message.Header.SrcAddr {
		t.Errorf("Clone Message failed.")
	}
	if newMessage.Body.Items[0] != message.Body.Items[0] {
		t.Errorf("Clone Message failed.")
	}
}
