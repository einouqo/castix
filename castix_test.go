package cool

import (
    "context"
    "sync"
    "testing"
)

func TestCastix(t *testing.T) {
    t.Run("notify", func(t *testing.T) {
        x := New[byte]()

        msgs := []byte{1, 2, 3}

        tests := []struct {
            name  string
            count int
        }{
            {name: "no one", count: 0},
            {name: "one", count: 1},
            {name: "few", count: 9},
            {name: "more", count: 10},
        }

        for _, test := range tests {
            t.Run(test.name, func(t *testing.T) {
                chs := make(map[<-chan byte]context.CancelFunc)

                for i := 0; i < test.count; i++ {
                    ctx := context.Background()
                    ctx, cancel := context.WithCancel(ctx)
                    chs[x.C(ctx)] = cancel
                }

                var wg sync.WaitGroup

                wg.Add(1)
                go func() {
                    defer wg.Done()
                    for _, msg := range msgs {
                        x.Notify(msg)
                    }
                }()

                for ch, cancel := range chs {
                    wg.Add(1)
                    go func(ch <-chan byte, cancel context.CancelFunc) {
                        defer wg.Done()
                        defer cancel()
                        for _, msg := range msgs {
                            r, ok := <-ch
                            if !ok {
                                t.Error("channel is closed")
                            }
                            if r != msg {
                                t.Errorf("received wrong message (want: %d, got: %d)", msg, r)
                            }
                        }
                    }(ch, cancel)
                }

                wg.Wait()
                x.Notify(0)
                if c := len(x.subs); c != 0 {
                    t.Errorf("inspection failed (expect 0 subscribers, got %d)", c)
                }
            })
        }

    })

    t.Run("drain", func(t *testing.T) {
        x := New[byte]()

        drain, last := []byte{1, 2, 3}, byte(4)

        ctx := context.Background()
        ctx, cancel := context.WithCancel(ctx)
        ch := x.C(ctx, WithDrain())

        for _, msg := range drain {
            x.Notify(msg)
        }
        x.Notify(last)

        cancel()
        x.Notify(0)

        if msg := <-ch; msg != last {
            t.Errorf("last message should be %d but got %d", last, msg)
        }
        if _, ok := <-ch; ok {
            t.Errorf("subscribe channel should be closed after cancel and read last")
        }
    })
}
