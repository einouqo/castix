package emit

import "github.com/einouqo/castix/internal/bridge"

type Output[T any] struct{ emitter[T] }

var _ bridge.Output[struct{}] = (*Output[struct{}])(nil)

func (o *Output[T]) init() *Output[T] {
	o.emitter.init()
	return o
}

func NewOutput[T any]() *Output[T] {
	return new(Output[T]).init()
}

func (o *Output[T]) Pass(v T) { o.emit(v) }

func (o *Output[T]) Watch(options ...bridge.OutputWatchOption) (<-chan T, bridge.Leave) {
	opts := make([]option, 0, len(options))
	for _, opt := range options {
		switch opt := opt.(type) {
		case bridge.WatchBufferSizeOption:
			opts = append(opts, withBuffSize(int(opt)))
		case bridge.WatchSkipOption:
			opts = append(opts, use(skipping))
		case bridge.WatchDrainOption:
			opts = append(opts, use(draining))
		case bridge.WatchFilterOption[T]:
			opts = append(opts, filtrate[T](filter[T](opt)))
		}
	}
	return o.watch(opts...)
}
