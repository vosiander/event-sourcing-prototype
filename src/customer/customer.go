package customer

import "eventsourcing"

type Customer struct {
	firstname string
	lastname  string
	createdAt string
	address   string
}

func Replay(es *eventsourcing.EventStream) *Customer {
	c := Customer{}
	stream := es.Stream()

	for _, e := range stream {
		switch true {
		case e.Name() == "Event.Create.Customer":
			c.firstname = e.Payload()["firstname"].(string)
			c.lastname = e.Payload()["lastname"].(string)
			c.createdAt = e.Payload()["createdAt"].(string)
			break
		case e.Name() == "Event.Update.Lastname":
			c.lastname = e.Payload()["lastname"].(string)
			break
		case e.Name() == "Event.Update.Address":
			c.address = e.Payload()["address"].(string)
			break
		}
	}

	return &c
}
