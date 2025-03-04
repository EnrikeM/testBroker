package main

import (
	"TestBroker/internal/httpsrv"
	"TestBroker/internal/queue"
	"log"
)

func ptr(s string) *string {
	return &s
}

var ExampleQueue1 = queue.Queue{
	Name: "queue1",
	Messages: []queue.Message{
		{
			EventType: ptr("queue1"),
		},
	},
	Capacity: 3,
}

var ExampleQueue2 = queue.Queue{
	Name: "queue2",
	Messages: []queue.Message{
		{
			EventType: nil,
		},
	},
	Capacity: 2,
}

var ExampleQueues = queue.Queues{
	Queues: []queue.Queue{
		ExampleQueue1,
		ExampleQueue2,
	},
}

func main() {

	api := httpsrv.NewAPI()
	api.SetQueues(ExampleQueues)
	err := api.Start()
	if err != nil {
		log.Fatal(err)
	}
}
