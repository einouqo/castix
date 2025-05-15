package emit

type strategy uint8

const (
	_ strategy = iota
	draining
	skipping
)

type config struct {
	size     int
	strategy strategy
}

func (c *config) init() *config {
	c.size = 1
	return c
}

type filter[T any] func(T) bool

type settings[T any] struct {
	config
	filter[T]
}

func (s *settings[T]) init() *settings[T] {
	s.config.init()
	return s
}

func build[T any](done <-chan struct{}, opts ...option) (chan T, deliver[T]) {
	s := new(settings[T]).init()
	for _, opt := range opts {
		switch opt := opt.(type) {
		case configure:
			opt(&s.config)
		case setup[T]:
			opt(s)
		}
	}

	ch, dlv := make(chan T, s.size), send[T](done)

	switch s.strategy {
	case draining:
		dlv = drain[T](done)
	case skipping:
		dlv = skip[T](done)
	}

	return ch, func(ch chan T, msg T) {
		if s.filter != nil && !s.filter(msg) {
			return
		}
		dlv(ch, msg)
	}
}
