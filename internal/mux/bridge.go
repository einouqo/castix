package mux

import "github.com/einouqo/castix/internal/bridge"

type Input[T any] struct{ mux[T] }

var _ bridge.Input[struct{}] = (*Input[struct{}])(nil)

func (in *Input[T]) init() *Input[T] {
	in.mux.init()
	return in
}

func NewInput[T any]() *Input[T] {
	return new(Input[T]).init()
}

func (in *Input[T]) Attach(ch <-chan T, h bridge.Handle[T], opts ...bridge.InputAttachOption) bridge.Leave {
	var f bridge.Filter[T]
	for _, opt := range opts {
		switch opt := opt.(type) {
		case bridge.AttachFilterOption[T]:
			f = bridge.Filter[T](opt)
		}
	}

	return in.attach(ch, func(msg T) {
		if f != nil && !f(msg) {
			return
		}
		h(msg)
	})
}
