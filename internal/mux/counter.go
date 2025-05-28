package mux

import "sync/atomic"

type counter struct{ v int32 }

func (c *counter) add(delta int) int {
	v := atomic.AddInt32(&c.v, int32(delta))
	return int(v)
}

func (c *counter) value() int {
	v := atomic.LoadInt32(&c.v)
	return int(v)
}
