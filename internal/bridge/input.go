package bridge

type Input[T any] interface {
	Attach(<-chan T, Handle[T], ...InputAttachOption) Leave
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
	_ InputAttachOption = (*AttachFilterOption[struct{}])(nil)
)
