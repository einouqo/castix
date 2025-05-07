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

func build[T any](done chan struct{}, opts ...option) (chan T, deliver[T]) {
  var (
    f   filter[T]
    cfg = new(config).init()
  )
  for _, opt := range opts {
    switch opt := opt.(type) {
    case configue:
      opt(cfg)
    case withFilter[T]:
      f = filter[T](opt)
    }
  }

  ch, dlv := make(chan T, cfg.size), send[T](done)

  switch cfg.strategy {
  case draining:
    dlv = drain[T](done)
  case skipping:
    dlv = skip[T](done)
  }

  return ch, func(ch chan T, msg T) {
    if f != nil && !f(msg) {
      return
    }
    dlv(ch, msg)
  }
}
