package emit

import (
	"sync"
	"sync/atomic"
)

var cil = (chan struct{})(nil)

type emitter[T any] struct {
	mu   sync.RWMutex
	outs map[chan T]deliver[T]
}

func (e *emitter[T]) init() *emitter[T] {
	e.outs = make(map[chan T]deliver[T])
	return e
}

func (e *emitter[T]) emit(v T) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	for out, dlv := range e.outs {
		dlv(out, v)
	}
}

func (e *emitter[T]) watch(opts ...option) (c <-chan T, stop func()) {
	done := atomic.Value{}
	done.Store(make(chan struct{}))

	ch, dlv := build[T](
		done.Load().(chan struct{}),
		opts...,
	)

	e.mu.Lock()
	e.outs[ch] = dlv
	e.mu.Unlock()

	return ch, func() {
		d := done.Swap(cil).(chan struct{})
		if d == cil {
			return // already stopped
		}
		close(d)

		e.mu.Lock()
		defer e.mu.Unlock()
		if _, hit := e.outs[ch]; hit {
			close(ch)
			delete(e.outs, ch)
		}
	}
}
