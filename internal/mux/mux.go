package mux

type detach = func()

type mux[T any] interface {
	attach(<-chan T, handle[T]) detach
	len() int
}

func create[T any](opts ...option) mux[T] {
	cfg := new(config).init()
	for _, opt := range opts {
		opt(cfg)
	}

	var x mux[T]
	switch cfg.variant {
	case goroutine:
		x = new(gux[T])
	case reflection:
		x = new(rux[T]).init()
	}

	return x
}
