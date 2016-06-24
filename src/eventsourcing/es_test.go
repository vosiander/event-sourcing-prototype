package eventsourcing

import (
	"testing"
)

func TestAddEvent(t *testing.T) {
	es := NewStream()
	es.Add(NewEvent("New test event"))
	es.Add(NewEvent("Second test event"))

	if !es.HasEvents() {
		t.Fail()
	}

	if es.Count() != 2 {
		t.Fail()
	}
}
