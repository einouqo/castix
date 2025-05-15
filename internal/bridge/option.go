package bridge

type AttachOption interface {
	InputAttachOption
	itsAttachOption()
}

var _ AttachOption = (*AttachFilterOption[struct{}])(nil)

type WatchOption interface {
	OutputWatchOption
	itsWatchOption()
}

var (
	_ WatchOption = (*WatchBufferSizeOption)(nil)
	_ WatchOption = (*WatchDrainOption)(nil)
	_ WatchOption = (*WatchSkipOption)(nil)
	_ WatchOption = (*WatchFilterOption[struct{}])(nil)
)
