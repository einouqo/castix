package bridge

type Leave = func()

type Convert[IN, OUT any] func(IN) OUT

type Handle[T any] func(T)

type Filter[T any] func(T) bool
