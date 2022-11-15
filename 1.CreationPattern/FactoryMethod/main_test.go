package factoryMethod_test

import (
	"testing"

	event "patterns/1.CreationPattern/FactoryMethod"
)

func TestFactory_Create(t *testing.T) {
	factory := new(event.Factory)
	start := factory.Create(event.Start)
	if start.EventType() != event.Start {
		t.Errorf("expect event.Start, but actual %v.", start.EventType())
	}
}
