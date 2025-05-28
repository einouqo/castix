package castix

import (
	"github.com/einouqo/castix/internal/bridge"
	"github.com/einouqo/castix/internal/mux"
)

type config struct {
	input []mux.InputOption
}

func (c *config) init() *config {
	c.input = make([]mux.InputOption, 0)
	return c
}

// Option defines a configuration option that can be applied when creating
// a new Castix instance using New.
type Option interface {
	apply(*config)
}

// UseGoroutine returns an Option that configures Castix to use goroutine-based
// multiplexing for handling its input channels. This is generally the default behavior
// if no specific input handling option is provided. In this mode, each attached
// input channel is monitored by a dedicated goroutine.
func UseGoroutine() Option {
	return funcOption{f: func(cfg *config) {
		cfg.input = append(cfg.input, mux.InputUseGoroutine())
	}}
}

// UseReflection returns an Option that configures Castix to use reflection-based
// multiplexing (via reflect.Select) for handling its input channels.
// This approach is recommended when dealing with a large number of input channels
// that are expected to receive messages infrequently. It is particularly well-suited
// for scenarios where conserving goroutine resources across many potentially idle
// channels is beneficial, as it uses a single goroutine to manage all inputs.
// For scenarios with fewer channels experiencing high-throughput and frequent messages,
// goroutine-based multiplexing (see UseGoroutine) may offer better performance.
func UseReflection() Option {
	return funcOption{f: func(cfg *config) {
		cfg.input = append(cfg.input, mux.InputUseReflection())
	}}
}

type funcOption struct {
	f func(*config)
}

var _ Option = (*funcOption)(nil)

func (o funcOption) apply(cfg *config) { o.f(cfg) }

// Filter defines a function signature for a predicate that determines
// whether a message of type T should be processed or discarded/skipped.
// Must return true if a trial is passed and false otherwise
type Filter[T any] func(T) bool

var _ = bridge.Filter[struct{}](Filter[struct{}](nil))

// SourceOption represents an option that can be applied when attaching a source
// to a Castix instance via the Source method.
type SourceOption = bridge.AttachOption

// WithSourceFilter returns a SourceOption that applies a filter to messages
// coming from a specific source. Only messages for which the filter function `f`
// returns true will be passed on for conversion and distribution.
// The filter is applied before type conversion.
func WithSourceFilter[T any](f Filter[T]) bridge.AttachFilterOption[T] {
	return bridge.WithAttachFilter[T](bridge.Filter[T](f))
}

// SubscribeOption represents an option that can be applied when creating a subscription
// to a Castix instance via the Subscribe method.
type SubscribeOption = bridge.WatchOption

var (
	// WithSubscribeBufferSize returns a SubscribeOption that configures the buffer
	// size of the channel returned by the Subscribe method.
	// A larger buffer can help prevent blocking (when using default delivery strategy)
	// if the subscriber processes messages slower than they are produced but consumes
	// more memory.
	WithSubscribeBufferSize = bridge.WithWatchBuffSize

	// WithSubscribeDrain returns a SubscribeOption that causes the subscription to
	// drop older messages in its buffer when it is full, removing the oldest message
	// to make room for the newest one. This strategy prioritizes newer messages at the
	// expense of potentially losing older ones when the subscriber cannot keep up with
	// the message rate.
	WithSubscribeDrain = bridge.WithWatchDrain

	// WithSubscribeSkip returns a SubscribeOption that causes the subscription to
	// drop messages if its buffer is full, instead of blocking the Castix multiplexer.
	// This can be useful when timely processing is less critical than overall system
	// responsiveness, but may lead to message loss.
	WithSubscribeSkip = bridge.WithWatchSkip
)

// WithSubscribeFilter returns a SubscribeOption that applies a filter to messages
// designated for a specific subscription. Only messages for which the filter function `f`
// returns true will be sent to the subscription's channel.
// The filter is applied before sending a message to the subscribed channel.
func WithSubscribeFilter[T any](f Filter[T]) bridge.WatchFilterOption[T] {
	return bridge.WithWatchFilter[T](bridge.Filter[T](f))
}
