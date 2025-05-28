package mux

import (
	"reflect"
	"sync"
)

type index = int

const (
	refresh index = iota
)

var indices = map[index]struct{}{
	refresh: {},
}

type handle[T any] func(T)

type rux[T any] struct {
	wakeup, refresh chan struct{}

	mu  sync.RWMutex
	ins map[<-chan T]handle[T]
}

var _ mux[struct{}] = (*rux[struct{}])(nil)

func (x *rux[T]) init() *rux[T] {
	x.wakeup = make(chan struct{}, 1)
	x.refresh = make(chan struct{})

	x.ins = make(map[<-chan T]handle[T])

	return x
}

func (x *rux[T]) attach(ch <-chan T, h handle[T]) detach {
	defer x.start()

	x.mu.Lock()
	x.ins[ch] = h
	x.mu.Unlock()

	return func() {
		defer x.start()

		x.mu.Lock()
		delete(x.ins, ch)
		x.mu.Unlock()
	}
}

func (x *rux[T]) len() int {
	x.mu.RLock()
	defer x.mu.RUnlock()
	return len(x.ins)
}

func (x *rux[T]) start() {
	select {
	case x.refresh <- struct{}{}:
	case x.wakeup <- struct{}{}:
		go func() {
			defer func() { <-x.wakeup }()
			x.loop()
		}()
	}
}

func (x *rux[T]) loop() {
	cases := []reflect.SelectCase{
		refresh: {
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(x.refresh),
		},
	}

REFRESH:
	for i := len(indices); i < len(cases); i++ {
		// prevent a potential memory leak: internal array may keep a pointer
		// to a channel that is no longer in use by the program
		cases[i] = reflect.SelectCase{}
	}
	cases = cases[:len(indices)]
	x.mu.RLock()
	for ch := range x.ins {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	x.mu.RUnlock()

	for {
		if len(cases) <= len(indices) {
			return
		}

		i, rv, ok := reflect.Select(cases)
		switch i {
		case refresh:
			goto REFRESH
		}

		ch := cases[i].Chan.Interface().(<-chan T)

		switch {
		case !ok: // the channel is closed
			l := len(cases) - 1
			cases[i], cases[l] = cases[l], reflect.SelectCase{}
			cases = cases[:l]
			x.mu.Lock()
			delete(x.ins, ch)
			x.mu.Unlock()
		default:
			x.mu.RLock()
			h, hit := x.ins[ch]
			x.mu.RUnlock()
			if hit {
				h(rv.Interface().(T))
			}
		}
	}
}
