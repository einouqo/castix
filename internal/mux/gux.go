package mux

import "sync/atomic"

var cil = (chan struct{})(nil)

type gux[T any] struct {
	count counter
}

var _ mux[struct{}] = (*gux[struct{}])(nil)

func (x *gux[T]) attach(ch <-chan T, h handle[T]) detach {
	stop := atomic.Value{}
	stop.Store(make(chan struct{}))

	x.count.add(1)
	done := make(chan struct{})
	go func(stop <-chan struct{}) {
		defer close(done)
		defer func() { x.count.add(-1) }()
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					return
				}
				h(v)
			case <-stop:
				return
			}
		}
	}(stop.Load().(chan struct{}))

	return func() {
		sp := stop.Swap(cil).(chan struct{})
		if sp == cil {
			return // already detached
		}
		close(sp)
		<-done
	}
}

func (x *gux[T]) len() int {
	return x.count.value()
}
