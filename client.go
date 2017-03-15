package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// SomeData is a dummy object to send via rpc
type SomeData struct {
	Schema    string `json:"$schema"`
	Timestamp string `json:"title"`
	Type      string `json:"type"`
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}

	someData := &SomeData{
		Schema:    "https://johnmccabe.net/schemas/go-rpc-test/draft-01/schema#1",
		Timestamp: time.Now().String(),
		Type:      "dummydata",
	}

	var ack bool
	err = client.Call("Scheduler.SubmitCapabilities", someData, &ack)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Ack received: %t", ack)
	}
}
