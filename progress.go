package maxconc

import (
	"sync"
)

type ChangeNotifier[T any] struct {
	mutex   sync.Mutex
	closed  bool
	last    T
	updates chan T

	subscribers []chan T
}

func (p *ChangeNotifier[T]) Last() T {
	return p.last
}

func (p *ChangeNotifier[T]) Subscribe() <-chan T {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return nil
	}

	c := make(chan T, 0)
	p.subscribers = append(p.subscribers, c)

	return c

}

func (p *ChangeNotifier[T]) Update(updates ...T) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return
	}

	for _, update := range updates {
		for _, sub := range p.subscribers {
			if len(sub) == 0 {
				sub <- update
			} else {
				drain(sub)
				sub <- update
			}
		}
	}

	p.last = updates[len(updates)-1]

}

func (p *ChangeNotifier[T]) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return
	}

	for _, sub := range p.subscribers {
		if len(sub) == 1 { // give a chance to slower readers to read
			drain(sub)
			close(sub)
		}

		// close everything
		//drain the channel
		for range sub {
		}
		close(sub)
	}
}

func drain[T any](c chan T) {
	//drain the channel
	for range c {
	}
}
