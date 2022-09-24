package fan

type Fan struct {
	Publish     chan string
	Subscribe   chan chan string
	Unsubscribe chan chan string

	subscribers map[*chan string]chan string
}

func (f *Fan) run() {
	for {
		select {
		case ch := <-f.Subscribe:
			f.subscribers[&ch] = ch
		case ch := <-f.Unsubscribe:
			delete(f.subscribers, &ch)
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
		subscribers: make(map[*chan string]chan string),
	}

	go f.run()

	return f
}
