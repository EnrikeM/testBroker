package config

import "TestBroker/internal/queue"

type Config struct {
	Queues []queue.Queue
}

func NewConfig(queues []queue.Queue) *Config {
	return &Config{
		Queues: queues,
	}
}
