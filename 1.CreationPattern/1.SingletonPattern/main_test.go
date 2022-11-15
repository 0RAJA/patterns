package singletonPattern

import (
	"testing"
)

func TestInstance(t *testing.T) {
	msg0 := Instance().GetMsg()
	if msg0.Count != 0 {
		t.Errorf("expect msg count %d, but actual %d.", 0, msg0.Count)
	}
	msg0.Count = 1
	Instance().AddMsg(msg0)
	msg1 := Instance().GetMsg()
	if msg1.Count != 1 {
		t.Errorf("expect msg count %d, but actual %d.", 0, msg1.Count)
	}
}
