package bridge

type Output[T any] interface {
  Pass(T)
  Watch(...OutputWatchOption) (<-chan T, Leave)
}

type WatchBufferSizeOption int

func (WatchBufferSizeOption) itsOutputWatchOption() {}
func (WatchBufferSizeOption) itsWatchOption()       {}

func WithWatchBuffSize(size int) WatchBufferSizeOption {
  return WatchBufferSizeOption(size)
}

type WatchDrainOption struct{}

func (WatchDrainOption) itsOutputWatchOption() {}
func (WatchDrainOption) itsWatchOption()       {}

func WithWatchDrain() (opt WatchDrainOption) { return opt }

type WatchSkipOption struct{}

func (WatchSkipOption) itsOutputWatchOption() {}
func (WatchSkipOption) itsWatchOption()       {}

func WithWatchSkip() (opt WatchSkipOption) { return opt }

type WatchFilterOption[T any] Filter[T]

func (WatchFilterOption[T]) itsOutputWatchOption() {}
func (WatchFilterOption[T]) itsWatchOption()       {}

func WithWatchFilter[T any](f Filter[T]) WatchFilterOption[T] {
  return WatchFilterOption[T](f)
}

type OutputWatchOption interface {
  itsOutputWatchOption()
}

var (
  _ OutputWatchOption = (*WatchBufferSizeOption)(nil)
  _ OutputWatchOption = (*WatchDrainOption)(nil)
  _ OutputWatchOption = (*WatchSkipOption)(nil)
)
