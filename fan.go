package fan

import (
	"fmt"
)

type Fan struct {
	Publish     chan string
	Subscribe   chan chan string
	Unsubscribe chan chan string

	subscribers map[string]chan string
}

func (f *Fan) run() {
	for {
		select {
		case ch := <-f.Subscribe:
			id := fmt.Sprintf("%v", ch)
			f.subscribers[id] = ch
		case ch := <-f.Unsubscribe:
			id := fmt.Sprintf("%v", ch)
			delete(f.subscribers, id)
		case msg := <-f.Publish:
			for _, sub := range f.subscribers {
				select {
				case sub <- msg:
					// sent
				default:
					// channel full, skip
				}
			}
		}
	}
}

func Run(ch chan string) *Fan {
	f := &Fan{
		Publish:     ch,
		Subscribe:   make(chan chan string),
		Unsubscribe: make(chan chan string),
		subscribers: make(map[string]chan string),
	}

	go f.run()

	return f
}
