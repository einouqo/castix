package castix

import (
	"github.com/einouqo/castix/internal/bridge"
	"github.com/einouqo/castix/internal/emit"
	"github.com/einouqo/castix/internal/mux"
)

type Leave = bridge.Leave

type Convert[IN, OUT any] func(IN) OUT

var _ = bridge.Convert[struct{}, struct{}](Convert[struct{}, struct{}](nil))

func Pass[T any](in T) T { return in }

var (
	_ Convert[struct{}, struct{}] = Pass[struct{}]
)

type Castix[IN, OUT any] struct {
	bridge *bridge.Bridge[IN, OUT]
}

func New[IN, OUT any](cv Convert[IN, OUT]) *Castix[IN, OUT] {
	return &Castix[IN, OUT]{
		bridge: bridge.New[IN, OUT](
			mux.NewInput[IN](), emit.NewOutput[OUT](),
			bridge.Convert[IN, OUT](cv),
		),
	}
}

func (x Castix[IN, OUT]) Source(ch <-chan IN, opts ...SourceOption) Leave {
	return x.bridge.Attach(ch, opts...)
}

func (x Castix[IN, OUT]) Subscribe(opts ...SubscribeOption) (<-chan OUT, Leave) {
	return x.bridge.Watch(opts...)
}
