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
	send := send[T](done)

	return func(ch chan T, msg T) {
		// prioritize non-blocking send before drain entry
		// it avoids an unnecessary drain attempt if the channel has capacity
		select {
		case ch <- msg:
			return
		default:
		}

		select {
		case ch <- msg:
		case <-ch:
			send(ch, msg)
		}
	}
}
