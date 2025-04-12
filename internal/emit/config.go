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

func build[T any](done chan struct{}, opts ...option) (chan T, deliver[T]) {
  cfg := new(config).init()
  for _, opt := range opts {
    opt(cfg)
  }

  ch, dlv := make(chan T, cfg.size), send[T](done)

  switch cfg.strategy {
  case draining:
    dlv = drain[T](done)
  case skipping:
    dlv = skip[T](done)
  }

  return ch, dlv
}
