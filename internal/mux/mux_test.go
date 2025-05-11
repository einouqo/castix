package mux

import (
	"sync"
	"testing"
	"time"
)

func TestMux(t *testing.T) {
	t.Run("attach", func(t *testing.T) {
		tests := []struct {
			name  string
			count int
		}{
			{name: "no one", count: 0},
			{name: "one", count: 1},
			{name: "few", count: 1 << 2},
			{name: "more", count: 1 << 4},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				x := new(mux[struct{}]).init()

				dets := make([]func(), 0, test.count)
				for i := 0; i < test.count; i++ {
					ch := make(chan struct{})
					t.Cleanup(func() { close(ch) })

					det := x.attach(ch, func(struct{}) {})
					dets = append(dets, det)
				}

				if l := x.len(); l != test.count {
					t.Errorf(
						"at least one channel was not attached (want: %d, got: %d)",
						test.count, l,
					)
				}

				for _, det := range dets {
					det()
				}

				if l := x.len(); l != 0 {
					t.Errorf("some were not detached (got: %d)", l)
				}
			})
		}
	})

	t.Run("handle", func(t *testing.T) {
		tests := []struct {
			name   string
			ins    [][]int
			expect map[int]int
		}{
			{name: "no source", ins: [][]int{}, expect: make(map[int]int)},
			{
				name:   "one source unique",
				ins:    [][]int{{1, 2, 3}},
				expect: map[int]int{1: 1, 2: 1, 3: 1},
			},
			{
				name:   "one source repeat",
				ins:    [][]int{{1, 3, 3}},
				expect: map[int]int{1: 1, 3: 2},
			},
			{
				name:   "few unique sources",
				ins:    [][]int{{10, 20}, {30, 40}, {50}},
				expect: map[int]int{10: 1, 20: 1, 30: 1, 40: 1, 50: 1},
			},
			{
				name:   "few sources repeat across",
				ins:    [][]int{{10, 20}, {20, 30}, {10, 30, 40}},
				expect: map[int]int{10: 2, 20: 2, 30: 2, 40: 1},
			},
			{
				name:   "few sources repeat internal",
				ins:    [][]int{{10, 10, 20}, {30, 40, 40}, {50}},
				expect: map[int]int{10: 2, 20: 1, 30: 1, 40: 2, 50: 1},
			},
			{
				name:   "few sources mixed repeat",
				ins:    [][]int{{10, 10, 20}, {20, 30, 30}, {10, 40}},
				expect: map[int]int{10: 3, 20: 2, 30: 2, 40: 1},
			},
			{
				name:   "few sources with empty",
				ins:    [][]int{{10, 20}, {}, {30, 30}, {40}},
				expect: map[int]int{10: 1, 20: 1, 30: 2, 40: 1},
			},
			{
				name:   "single empty source",
				ins:    [][]int{{}},
				expect: make(map[int]int),
			},
			{
				name:   "all empty sources",
				ins:    [][]int{{}, {}, {}},
				expect: make(map[int]int),
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				x := new(mux[int]).init()

				total, rmp := 0, make(map[int]int, len(test.expect))
				for k, v := range test.expect {
					total += v
					rmp[k] = v
				}

				mu := sync.Mutex{}
				calls := 0
				h := func(s int) {
					mu.Lock()
					defer mu.Unlock()
					rmp[s]--
					calls++
				}

				wg := sync.WaitGroup{}
				for _, in := range test.ins {
					wg.Add(1)
					go func(in []int) {
						defer wg.Done()

						ch := make(chan int)
						defer close(ch)

						x.attach(ch, h)

						for _, s := range in {
							ch <- s
						}
					}(in)
				}
				wg.Wait()

				for x.len() != 0 {
					time.Sleep(time.Millisecond)
				}

				if calls != total {
					t.Errorf("less handle calls than expected (want: %d, got: %d)", total, calls)
				}
				for v, c := range rmp {
					if _, hit := test.expect[v]; !hit {
						t.Errorf("handled unexpected message (value: %d)", v)
					}
					if c != 0 {
						t.Errorf("wrong c of value handle (value: %d, count: %d)", v, c)
					}
				}
			})
		}
	})
}
