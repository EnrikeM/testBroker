package queue

import "fmt"

type Message struct {
	EventType *string `json:"event_type"`
}

type Queues struct {
	Queues []Queue
}

type Queue struct {
	Name     string
	Messages []Message
	Capacity int
}

func NewQueue(name string, capacity int, messages []Message) *Queue {
	return &Queue{
		Name:     name,
		Capacity: capacity,
		Messages: messages,
	}
}

func (q *Queues) SetMessage(receivedMsg Message, queueName string) error {
	if receivedMsg.EventType == nil {
		return fmt.Errorf("empty message")
	}
	// идём по очереди и находим свободный слот для сообщения
	var msgSet bool

	for _, queue := range q.Queues {
		if msgSet {
			break
		}
		if len(queue.Messages) == queue.Capacity {
			return fmt.Errorf("queue is full")
		}

		for i, msg := range queue.Messages {
			if msg.EventType == nil {
				queue.Messages[i] = receivedMsg
				msgSet = true
				break
			}
		}
	}
	if !msgSet {
		return fmt.Errorf("err setting message")
	}

	return nil
}
