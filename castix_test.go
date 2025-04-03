package castix

import (
    "sync"
    "testing"
)

func FuzzCastixCancel(f *testing.F) {
    f.Add(1)

    f.Fuzz(func(t *testing.T, count int) {
        x := New[byte]()
        ch, cancel := x.C()
        cancel()

        for i := 0; i < count; i++ {
            cancel()
        }

        if _, open := <-ch; open {
            t.Fatal("channel is still open")
        }
    })
}

func TestCastix(t *testing.T) {
    t.Run("safe cancel", func(t *testing.T) {
        x := New[byte]()
        _, cancel := x.C()
        cancel()
        cancel()
    })

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
                chs := make(map[<-chan byte]Cancel)

                for i := 0; i < test.count; i++ {
                    ch, cancel := x.C()
                    chs[ch] = cancel
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
                    go func(ch <-chan byte, cancel Cancel) {
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
                if c := len(x.subs); c != 0 {
                    t.Errorf("inspection failed (expect 0 subscribers, got %d)", c)
                }
            })
        }

    })

    t.Run("drain", func(t *testing.T) {
        x := New[byte]()

        drain, last := []byte{1, 2, 3}, byte(4)

        ch, cancel := x.C(WithDrain())

        for _, msg := range drain {
            x.Notify(msg)
        }
        x.Notify(last)

        cancel()
        if msg := <-ch; msg != last {
            t.Errorf("last message should be %d but got %d", last, msg)
        }
        if _, ok := <-ch; ok {
            t.Errorf("subscribe channel should be closed after cancel and read last")
        }
    })
}

func BenchmarkCastix(b *testing.B) {
    b.Run("notify", func(b *testing.B) {
        benches := []struct {
            name  string
            count int
        }{
            {name: "no one", count: 0},
            {name: "one", count: 1},
            {name: "few", count: 1 << 2},
            {name: "more", count: 1 << 4},
            {name: "even more", count: 1 << 8},
            {name: "a lot", count: 1 << 16},
        }

        for _, bench := range benches {
            x := New[int]()
            cancels := make([]Cancel, 0, bench.count)

            var wg sync.WaitGroup
            for i := 0; i < bench.count; i++ {
                ch, cancel := x.C()
                cancels = append(cancels, cancel)

                wg.Add(1)
                go func(ch <-chan int) {
                    defer wg.Done()
                    for range ch {
                    }
                }(ch)
            }

            b.Run(bench.name, func(b *testing.B) {
                for i := 0; i < b.N; i++ {
                    x.Notify(i)
                }
            })

            for _, cancel := range cancels {
                cancel()
            }
            wg.Wait()
        }
    })

    b.Run("rapid notify", func(b *testing.B) {
        b.Skip("implement me")
    })

    b.Run("runtime ops", func(b *testing.B) {
        b.Skip("implement me")
    })
}
