package bridge

type Input[T any] interface {
  Attach(<-chan T, ...InputAttachOption) Leave
}

type AttachHandleOption[T any] Handle[T]

func (AttachHandleOption[T]) itsInputAttachOption() {}

func WithAttachHandle[T any](handle Handle[T]) AttachHandleOption[T] {
  return AttachHandleOption[T](handle)
}

type AttachFilterOption[T any] Filter[T]

func WithAttachFilter[T any](f Filter[T]) AttachFilterOption[T] {
  return AttachFilterOption[T](f)
}

func (AttachFilterOption[T]) itsInputAttachOption() {}
func (AttachFilterOption[T]) itsAttachOption()      {}

type InputAttachOption interface {
  itsInputAttachOption()
}

var (
  _ InputAttachOption = (*AttachHandleOption[struct{}])(nil)
  _ InputAttachOption = (*AttachFilterOption[struct{}])(nil)
)
