package mux

import (
	"testing"
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
				mux := new(mux[struct{}]).init()

				dets := make([]func(), 0, test.count)
				for i := 0; i < test.count; i++ {
					ch := make(chan struct{})
					t.Cleanup(func() { close(ch) })

					det := mux.attach(ch, func(struct{}) {})
					dets = append(dets, det)
				}

				if mux.len() != test.count {
					t.Errorf(
						"at least one channel was not attached (want: %d, got: %d)",
						test.count, len(mux.ins),
					)
				}

				for _, det := range dets {
					det()
				}

				if mux.len() != 0 {
					t.Errorf("some were not detached (got: %d)", len(mux.ins))
				}
			})
		}
	})

	t.Run("handle", func(t *testing.T) {
		t.Skip("implement me")
	})
}
