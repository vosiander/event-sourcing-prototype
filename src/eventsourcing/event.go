package eventsourcing

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"log"
)

type Event struct {
	id      uuid.UUID
	payload map[string]interface{}
}

func NewEvent(s string) Event {
	payload := []byte(s)
	var f map[string]interface{}
	e := Event{id: uuid.NewV4()}

	json.Unmarshal(payload, &f)
	log.Println(f)
	e.payload = f

	return e
}

func (e *Event) Payload() map[string]interface{} {
	return e.payload
}

func (e *Event) Name() string {
	return e.payload["name"].(string)
}
