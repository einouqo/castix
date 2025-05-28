package mux

type variant uint8

const (
	goroutine variant = iota + 1
	reflection
)

type config struct {
	variant variant
}

func (c *config) init() *config {
	c.variant = goroutine
	return c
}

type option func(*config)

func use(v variant) option {
	return func(cfg *config) { cfg.variant = v }
}
