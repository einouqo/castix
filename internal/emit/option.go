package emit

type configue func(*config)

var (
  _ option = configue(nil)
)

func (configue) itsOption() {}

func withBuffSize(size int) configue {
  return func(c *config) {
    if size >= 1 {
      c.size = size
    }
  }
}

func withStrategy(s strategy) configue {
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
