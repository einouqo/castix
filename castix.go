package castix

import "sync"

type control struct {
    drain bool
}

type Castix[T any] struct {
    mu   sync.RWMutex
    subs map[chan T]control
}

func New[T any]() *Castix[T] {
    return &Castix[T]{
        subs: make(map[chan T]control),
    }
}

type Cancel func()

func (x *Castix[T]) C(opts ...Option) (<-chan T, Cancel) {
    ch := make(chan T, 1)
    ctr := control{}
    for _, opt := range opts {
        opt.apply(&ctr)
    }

    x.mu.Lock()
    x.subs[ch] = ctr
    x.mu.Unlock()

    return ch, func() {
        x.mu.Lock()
        defer x.mu.Unlock()
        if _, hit := x.subs[ch]; hit {
            close(ch)
            delete(x.subs, ch)
        }
    }
}

func (x *Castix[T]) Notify(msg T) {
    x.mu.RLock()
    defer x.mu.RUnlock()
    for ch, ctr := range x.subs {
        if !ctr.drain {
            ch <- msg
            continue
        }

        select {
        case ch <- msg:
            continue
        default:
        }

        select {
        case ch <- msg:
        case <-ch:
            ch <- msg
        }
    }
}
