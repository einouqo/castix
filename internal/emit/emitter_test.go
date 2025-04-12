package emit

import (
  "sync"
  "testing"
)

func TestEmitter(t *testing.T) {
  t.Run("emit", func(t *testing.T) {
    p := new(emitter[byte]).init()

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
        chs := make(map[<-chan byte]func())

        for i := 0; i < test.count; i++ {
          ch, stop := p.watch()
          chs[ch] = stop
        }

        var wg sync.WaitGroup

        wg.Add(1)
        go func() {
          defer wg.Done()
          for _, msg := range msgs {
            p.emit(msg)
          }
        }()

        for ch, cancel := range chs {
          wg.Add(1)
          go func(ch <-chan byte, stop func()) {
            defer wg.Done()
            defer stop()
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
        if c := len(p.outs); c != 0 {
          t.Errorf("inspection failed (expect 0 subscribers, got %d)", c)
        }
      })
    }
  })

  t.Run("strategy", func(t *testing.T) {
    p := new(emitter[byte]).init()

    msgs := []byte{1, 2, 3}

    tests := []struct {
      name     string
      strategy strategy
      expect   byte
    }{
      {name: "skipping", strategy: skipping, expect: 1},
      {name: "draining", strategy: draining, expect: 3},
    }

    for _, test := range tests {
      t.Run(test.name, func(t *testing.T) {
        ch, stop := p.watch(withStrategy(test.strategy))

        for _, msg := range msgs {
          p.emit(msg)
        }

        stop()
        if msg := <-ch; msg != test.expect {
          t.Errorf("message should be %d but got %d", test.expect, msg)
        }
        if _, ok := <-ch; ok {
          t.Errorf("subscribe channel should be closed after stop and read last")
        }
      })
    }
  })
}
