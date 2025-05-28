package mux

import "github.com/einouqo/castix/internal/bridge"

type InputOption option

func InputUseGoroutine() InputOption {
	return InputOption(use(goroutine))
}

func InputUseReflection() InputOption {
	return InputOption(use(reflection))
}

type Input[T any] struct{ x mux[T] }

var _ bridge.Input[struct{}] = (*Input[struct{}])(nil)

func NewInput[T any](options ...InputOption) *Input[T] {
	opts := make([]option, 0, len(options))
	for _, opt := range options {
		opts = append(opts, option(opt))
	}
	return &Input[T]{x: create[T](opts...)}
}

func (in *Input[T]) Attach(ch <-chan T, h bridge.Handle[T], opts ...bridge.InputAttachOption) bridge.Leave {
	var f bridge.Filter[T]
	for _, opt := range opts {
		switch opt := opt.(type) {
		case bridge.AttachFilterOption[T]:
			f = bridge.Filter[T](opt)
		}
	}

	return in.x.attach(ch, func(v T) {
		if f != nil && !f(v) {
			return
		}
		h(v)
	})
}
