package main

import (
	"fmt"
	"time"
)

type Event struct {
	queue   chan int      // channel
	timeout time.Duration // the timeout of produce send data
	handle  func(v int)   // consume callback
}

// NewEvent new queue event
func NewEvent(buffer int, t time.Duration, f func(v int)) *Event {
	return &Event{
		queue:   make(chan int, buffer),
		timeout: t,
		handle:  f,
	}
}

// Producer produce to queue
func (e Event) Producer(num int) {
	select {
	case e.queue <- num:
	case <-time.After(e.timeout):
	}
}

// Consumer consume from queue
func (e Event) Consumer() {
	for v := range e.queue {
		e.handle(v)
	}
}

func main() {
	e := NewEvent(10, time.Second, func(v int) {
		fmt.Println("this is num : ", v)
	})
	go func() {
		for i := 0; i < 10; i++ {
			e.Producer(i)
		}
	}()
	go func() {
		e.Consumer()
	}()
	time.Sleep(5 * time.Second)
}
