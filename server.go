package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

// SomeData is a dummy object to send via rpc
type SomeData struct {
	Schema    string `json:"$schema"`
	Timestamp string `json:"title"`
	Type      string `json:"type"`
}

type Ack bool

type Scheduler int

// SubmitCapabilities used to submit data to the server
func (l *Scheduler) SubmitCapabilities(data *SomeData, ack *Ack) error {
	fmt.Printf("Schema: %s\n", data.Schema)
	fmt.Printf("Title: %s\n", data.Timestamp)
	fmt.Printf("Type: %s\n", data.Type)
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Unable to unmarshal data received from client: %s\n", err)
	} else {
		fmt.Println(string(b))
		*ack = true
	}
	return nil
}

func f(from string) {
	for i := 0; i < 1000; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go f("goroutinetop")

	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	scheduler := new(Scheduler)
	rpc.Register(scheduler)
	rpc.Accept(inbound)
	go f("goroutinebottom")

}
