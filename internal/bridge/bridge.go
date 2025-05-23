package bridge

type Bridge[IN, OUT any] struct {
	in  Input[IN]
	out Output[OUT]

	cv Convert[IN, OUT]
}

func New[IN, OUT any](
	in Input[IN], out Output[OUT],
	cv Convert[IN, OUT],
) *Bridge[IN, OUT] {
	return &Bridge[IN, OUT]{
		in: in, out: out,
		cv: cv,
	}
}

func (b Bridge[IN, OUT]) Attach(ch <-chan IN, options ...AttachOption) Leave {
	opts := make([]InputAttachOption, 0, len(options)+1)
	for _, option := range options {
		opts = append(opts, option)
	}
	return b.in.Attach(ch, func(in IN) { b.out.Pass(b.cv(in)) }, opts...)
}

func (b Bridge[IN, OUT]) Watch(options ...WatchOption) (<-chan OUT, Leave) {
	opts := make([]OutputWatchOption, 0, len(options))
	for _, option := range options {
		opts = append(opts, option)
	}
	return b.out.Watch(opts...)
}
