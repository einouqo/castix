package castix

import "github.com/einouqo/castix/internal/bridge"

type Filter[T any] func(T) bool

var _ = bridge.Filter[struct{}](Filter[struct{}](nil))

type SourceOption = bridge.AttachOption

func WithSourceFilter[T any](f Filter[T]) bridge.AttachFilterOption[T] {
	return bridge.WithAttachFilter[T](bridge.Filter[T](f))
}

type SubscribeOption = bridge.WatchOption

var (
	WithSubscribeBufferSize = bridge.WithWatchBuffSize
	WithSubscribeDrain      = bridge.WithWatchDrain
	WithSubscribeSkip       = bridge.WithWatchSkip
)

func WithSubscribeFilter[T any](f Filter[T]) bridge.WatchFilterOption[T] {
	return bridge.WithWatchFilter[T](bridge.Filter[T](f))
}
