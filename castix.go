package cool

import (
    "context"
    "sync"
)

type control struct {
    done  <-chan struct{}
    drain bool
}

func (ctr *control) init(done <-chan struct{}) *control {
    ctr.done, ctr.drain = done, false
    return ctr
}

func (ctr *control) cleanup() {
    ctr.done = nil
}

type Castix[T any] struct {
    mu   sync.Mutex
    subs map[chan T]*control
}

func New[T any]() *Castix[T] {
    return &Castix[T]{
        subs: make(map[chan T]*control),
    }
}

func (x *Castix[T]) C(ctx context.Context, opts ...Option) <-chan T {
    ch := make(chan T, 1)
    ctr := new(control).init(ctx.Done())
    for _, opt := range opts {
        opt.apply(ctr)
    }

    x.mu.Lock()
    x.subs[ch] = ctr
    x.mu.Unlock()

    return ch
}

func (x *Castix[T]) Notify(msg T) {
    x.mu.Lock()
    defer x.mu.Unlock()
    for ch, ctr := range x.subs {
        select {
        case <-ctr.done:
            close(ch)
            ctr.cleanup()
            delete(x.subs, ch)
            continue
        default:
        }

        select {
        case ch <- msg:
            continue
        default:
        }

        if ctr.drain {
            <-ch
        }
        ch <- msg
    }
}
