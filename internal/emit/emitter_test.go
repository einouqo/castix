package emit

import (
	"sync"
	"testing"
)

func TestEmitter(t *testing.T) {
	t.Run("emit", func(t *testing.T) {
		e := new(emitter[byte]).init()

		vs := []byte{1, 2, 3}

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
					ch, stop := e.watch()
					chs[ch] = stop
				}

				var wg sync.WaitGroup

				wg.Add(1)
				go func() {
					defer wg.Done()
					for _, v := range vs {
						e.emit(v)
					}
				}()

				for ch, stop := range chs {
					wg.Add(1)
					go func(ch <-chan byte, stop func()) {
						defer wg.Done()
						defer stop()
						for _, v := range vs {
							r, ok := <-ch
							if !ok {
								t.Error("channel is closed")
							}
							if r != v {
								t.Errorf("received wrong message (want: %d, got: %d)", v, r)
							}
						}
					}(ch, stop)
				}

				wg.Wait()
				if c := len(e.outs); c != 0 {
					t.Errorf("inspection failed (expect 0 subscribers, got %d)", c)
				}
			})
		}
	})

	t.Run("strategy", func(t *testing.T) {
		e := new(emitter[byte]).init()

		vs := []byte{1, 2, 3}

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
				ch, stop := e.watch(use(test.strategy))

				for _, v := range vs {
					e.emit(v)
				}

				stop()
				if v := <-ch; v != test.expect {
					t.Errorf("message should be %d but got %d", test.expect, v)
				}
				if _, ok := <-ch; ok {
					t.Errorf("subscribe channel should be closed after stop and read last")
				}
			})
		}
	})

	t.Run("filter", func(t *testing.T) {
		e := new(emitter[int]).init()

		vs := []int{-3, -2, -1, 0, 1, 2, 3}

		tests := []struct {
			name   string
			filter filter[int]
			expect []int
		}{
			{name: "odds", filter: func(k int) bool { return k%2 != 0 }, expect: []int{-3, -1, 1, 3}},
			{name: "evens", filter: func(k int) bool { return k%2 == 0 }, expect: []int{-2, 0, 2}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				ch, stop := e.watch(filtrate(test.filter))
				defer stop()

				done := make(chan struct{})
				go func() {
					defer close(done)
					for _, v := range vs {
						e.emit(v)
					}
				}()

				for i, v := range test.expect {
					r, ok := <-ch
					if !ok {
						t.Error("channel is closed")
					}
					if r != v {
						t.Errorf("received wrong message (iteration: %d, want: %d, got: %d)", i, v, r)
					}
				}

				<-done
			})
		}
	})
}
