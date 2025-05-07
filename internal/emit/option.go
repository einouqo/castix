package emit

type configue func(*config)

var _ option = configue(nil)

func (configue) itsOption() {}

func withBuffSize(size int) configue {
  return func(c *config) {
    if size >= 1 {
      c.size = size
    }
  }
}

func withStrategy(s strategy) configue {
  return func(c *config) { c.strategy = s }
}

type withFilter[T any] filter[T]

var _ option = withFilter[struct{}](nil)

func (withFilter[T]) itsOption() {}

type option interface {
  itsOption()
}
