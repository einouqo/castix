package emit

type deliver[T any] func(chan T, T)

func send[T any](done <-chan struct{}) deliver[T] {
	return func(ch chan T, msg T) {
		select {
		case ch <- msg:
		case <-done:
		}
	}
}

func skip[T any](_ <-chan struct{}) deliver[T] {
	return func(ch chan T, msg T) {
		select {
		case ch <- msg:
		default:
		}
	}
}

func drain[T any](done <-chan struct{}) deliver[T] {
	return func(ch chan T, msg T) {
		select {
		case ch <- msg:
			return
		default:
		}

		select {
		case ch <- msg:
		case <-ch:
			select {
			case ch <- msg:
			case <-done:
			}
		}
	}
}
