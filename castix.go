package castix

import (
	"github.com/einouqo/castix/internal/bridge"
	"github.com/einouqo/castix/internal/emit"
	"github.com/einouqo/castix/internal/mux"
)

// Leave is a function that should be called to clean up resources
type Leave = bridge.Leave

// Convert defines a function signature for converting a value of type IN
// to a value of type OUT. This is used by Castix to transform messages
// from their source type to the desired output type for subscribers.
type Convert[IN, OUT any] func(IN) OUT

var _ = bridge.Convert[struct{}, struct{}](Convert[struct{}, struct{}](nil))

// Pass is a utility Convert function that returns the input value unchanged.
// It can be used when no type conversion is necessary between the source and subscriber (i.e., IN and OUT are the same type).
func Pass[T any](in T) T { return in }

var _ Convert[struct{}, struct{}] = Pass[struct{}]

// Castix provides a flexible and type-safe way to manage message streams.
// It acts as a multiplexer, allowing multiple message producers (sources)
// to broadcast messages to multiple consumers (subscribers).
// Castix also handles type conversion between the source message type (IN)
// and the subscriber's desired message type (OUT) using a provided Convert function.
type Castix[IN, OUT any] struct {
	bridge *bridge.Bridge[IN, OUT]
}

// New creates and initializes a new Castix instance.
// It requires a Convert function `cv` that defines how messages of type IN
// are transformed into messages of type OUT. If no transformation is needed
// (IN and OUT are the same type), the Pass function can be used.
func New[IN, OUT any](cv Convert[IN, OUT]) *Castix[IN, OUT] {
	return &Castix[IN, OUT]{
		bridge: bridge.New[IN, OUT](
			mux.NewInput[IN](), emit.NewOutput[OUT](),
			bridge.Convert[IN, OUT](cv),
		),
	}
}

// Source attaches a Go channel `ch` as a message producer to the Castix instance.
// Messages sent to `ch` will be processed by Castix, converted according to the
// function provided to New, and then distributed to all active subscribers.
//
// Callers can provide SourceOption arguments to customize the behavior of the source.
// Available options include:
//   - WithSourceFilter: Applies a filter to messages before they are sent to the subscribers.
//
// It returns a Leave function. This function MUST be called when the source channel
// is no longer needed. Calling Leave detaches the source from Castix and releases
// associated resources. The source will also be detached automatically if the provided
// `ch` channel is closed. In both scenarios, resources are cleaned up.
func (x Castix[IN, OUT]) Source(ch <-chan IN, opts ...SourceOption) Leave {
	return x.bridge.Attach(ch, opts...)
}

// Subscribe creates a new subscription to receive messages processed by the Castix instance.
// It returns two values:
//  1. A receive-only channel on which converted messages of type OUT
//     will be delivered.
//  2. A Leave function.
//
// Messages sent to this channel originate from all attached sources and are converted
// using the Convert function specified when the Castix instance was created.
//
// Callers can provide `SubscribeOption` arguments to customize the subscription's
// behavior. Available options include:
//   - WithSubscribeBufferSize: Configures the buffer size of the returned message channel.
//   - WithSubscribeFilter: Applies a filter to messages before they are sent to the channel.
//   - WithSubscribeDrain: Ensures all buffered messages are processed when unsubscribing.
//   - WithSubscribeSkip: Allows dropping messages if the subscription channel's buffer is full.
//
// The returned Leave function MUST be called when the subscription is no longer
// needed. This will close the message channel and release all associated resources.
// It's important to call Leave to prevent indefinite goroutine blocking and resource leaks.
func (x Castix[IN, OUT]) Subscribe(opts ...SubscribeOption) (<-chan OUT, Leave) {
	return x.bridge.Watch(opts...)
}
