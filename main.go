package main

import (
	"customer"
	"eventsourcing"
	"log"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	/*
		r := mux.NewRouter()
		r.HandleFunc("/", YourHandler)

		log.Fatal(http.ListenAndServe(":8000", r))
	*/
	es := eventsourcing.RecreateFromFile("./event_stream.txt")
	e := eventsourcing.NewEvent("{\"name\": \"Event.Update.Lastname\",\"lastname\": \"Petermann\"}")
	es.Add(e)
	es.Persist("./event_stream_output.txt")

	log.Println("Event Name: ", e.Name())

	c := customer.Replay(es)
	log.Println(c)
}
