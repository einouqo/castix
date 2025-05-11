package emit

type configure func(*config)

var (
  _ option = configure(nil)
)

func (configure) itsOption() {}

func withBuffSize(size int) configure {
  return func(c *config) {
    if size >= 1 {
      c.size = size
    }
  }
}

func withStrategy(s strategy) configure {
  return func(c *config) { c.strat = s }
}

type setup[T any] func(*settings[T])

var (
  _ option = setup[struct{}](nil)
)

func (s setup[T]) itsOption() {}

func withFilter[T any](f filter[T]) setup[T] {
  return func(s *settings[T]) { s.filter = f }
}

type option interface {
  itsOption()
}
