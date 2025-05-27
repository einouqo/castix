package emit

type configure func(*config)

var _ option = configure(nil)

func (configure) itsOption() {}

func use(s strategy) configure {
	return func(c *config) { c.strategy = s }
}
func withBuffSize(size int) configure {
	return func(c *config) {
		if size >= 1 {
			c.size = size
		}
	}
}

type setup[T any] func(*settings[T])

var _ option = setup[struct{}](nil)

func (s setup[T]) itsOption() {}

func filtrate[T any](f filter[T]) setup[T] {
	return func(s *settings[T]) { s.filter = f }
}

type option interface {
	itsOption()
}
