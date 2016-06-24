package eventsourcing

import "testing"

func TestNewEvent(t *testing.T) {
	e := NewEvent("{\"bla\": \"test\"}")

	if len(e.Payload()) == 0 {
		t.Fail()
	}
}
