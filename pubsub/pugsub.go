package main

import "sync"

type Broadcaster[T any] struct {
	mu   sync.Mutex
	subs map[chan T]struct{}
	last *T
}

func (b *Broadcaster[T]) Subscribe() <-chan T {
	ch := make(chan T, 1)

	b.mu.Lock()
	if b.last != nil {
		ch <- *b.last // catch-up
	}
	b.subs[ch] = struct{}{}
	b.mu.Unlock()

	return ch
}

func (b *Broadcaster[T]) Publish(v T) {
	b.mu.Lock()
	b.last = &v
	for ch := range b.subs {
		select {
		case ch <- v:
		default:
		}
	}
	b.mu.Unlock()
}
