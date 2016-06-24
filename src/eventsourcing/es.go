package eventsourcing

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type EventStream struct {
	stream []Event
}

func NewStream() *EventStream {
	var stream = []Event{}

	return &EventStream{stream: stream}
}

func Recreate(stream []Event) *EventStream {
	return &EventStream{stream: stream}
}

func RecreateFromFile(filePath string) *EventStream {
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := bufio.NewScanner(file)

	es := NewStream()

	for reader.Scan() {
		payload := reader.Text()

		if len(payload) > 0 {
			es.Add(NewEvent(payload))
		}
	}

	return es
}

func (es *EventStream) Persist(output string) {
	f, _ := os.Create(output)
	w := bufio.NewWriter(f)
	defer w.Flush()

	for i := range es.stream {
		s, _ := json.Marshal(es.stream[i].Payload())
		w.Write(s)
		w.WriteString("\n")
	}
}

func (es *EventStream) Stream() []Event {
	return es.stream
}

func (es *EventStream) Add(e Event) {
	log.Println(e)
	es.stream = append(es.stream, e)
}

func (es *EventStream) HasEvents() bool {
	return es.Count() > 0
}

func (es *EventStream) Count() int {
	return len(es.stream)
}
