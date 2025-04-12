package mux

import "github.com/einouqo/castix/internal/bridge"

type Input[T any] struct{ mux[T] }

var (
  _ bridge.Input[struct{}] = (*Input[struct{}])(nil)
)

func (in *Input[T]) init() *Input[T] {
  in.mux.init()
  return in
}

func NewInput[T any]() *Input[T] {
  return new(Input[T]).init()
}

func (in *Input[T]) Attach(ch <-chan T, options ...bridge.InputAttachOption) bridge.Leave {
  h := handle[T](func(T) {})
  for _, option := range options {
    switch opt := option.(type) {
    case bridge.AttachHandleOption[T]:
      h = handle[T](opt)
    case bridge.AttachFilterOption[T]:
      // TODO: implement me
      panic("implement me")
    }
  }
  return in.attach(ch, h)
}
