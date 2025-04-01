package castix

type Option interface {
	apply(*control)
}

func WithDrain() Option {
	return funcOption{
		fn: func(ctr *control) { ctr.drain = true },
	}
}

type funcOption struct {
	fn func(*control)
}

var _ Option = (*funcOption)(nil)

func (f funcOption) apply(ctr *control) {
	f.fn(ctr)
}
