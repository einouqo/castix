package emit

import "github.com/einouqo/castix/internal/bridge"

type Output[T any] struct{ emitter[T] }

var (
  _ bridge.Output[struct{}] = (*Output[struct{}])(nil)
)

func (o *Output[T]) init() *Output[T] {
  o.emitter.init()
  return o
}

func NewOutput[T any]() *Output[T] {
  return new(Output[T]).init()
}

func (o *Output[T]) Pass(msg T) { o.emit(msg) }

func (o *Output[T]) Watch(options ...bridge.OutputWatchOption) (<-chan T, bridge.Leave) {
  opts := make([]option, 0, len(options))
  for _, option := range options {
    switch opt := option.(type) {
    case bridge.WatchBufferSizeOption:
      opts = append(opts, withBuffSize(int(opt)))
    case bridge.WatchSkipOption:
      opts = append(opts, withStrategy(skipping))
    case bridge.WatchDrainOption:
      opts = append(opts, withStrategy(draining))
    case bridge.WatchFilterOption[T]:
      // TODO: implement me
      panic("implement me")
    }
  }
  return o.watch(opts...)
}
