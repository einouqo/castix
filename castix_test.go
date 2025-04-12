package castix

import (
	"context"
	"sync"
	"testing"
)

func FuzzCastixCancel(f *testing.F) {
	f.Add(1)

	f.Fuzz(func(t *testing.T, count int) {
		x := New[struct{}, struct{}](Pass[struct{}])
		ch, leave := x.Subscribe()
		leave()

		for i := 0; i < count; i++ {
			leave()
		}

		if _, open := <-ch; open {
			t.Fatal("channel is still open")
		}
	})
}

func BenchmarkCastix(b *testing.B) {
	b.Run("static", func(b *testing.B) {
		type run func(b *testing.B, n int)
		type runner func(b *testing.B, f run)

		runners := []runner{
			func(b *testing.B, f run) { b.Run("no one", func(b *testing.B) { f(b, 0) }) },
			func(b *testing.B, f run) { b.Run("one", func(b *testing.B) { f(b, 1) }) },
			func(b *testing.B, f run) { b.Run("few", func(b *testing.B) { f(b, 1<<2) }) },
			func(b *testing.B, f run) { b.Run("more", func(b *testing.B) { f(b, 1<<4) }) },
			func(b *testing.B, f run) { b.Run("even more", func(b *testing.B) { f(b, 1<<8) }) },
			func(b *testing.B, f run) { b.Run("a lot", func(b *testing.B) { f(b, 1<<10) }) },
		}

		b.Run("inputs", func(b *testing.B) {
			for _, rr := range runners {
				rr(b, func(b *testing.B, n int) {
					x := New[int, int](Pass[int])

					ctx := context.Background()
					ctx, cancel := context.WithCancel(ctx)

					wg := sync.WaitGroup{}

					for i := 0; i < n; i++ {
						wg.Add(1)
						go func() {
							defer wg.Done()

							in := make(chan int)
							defer close(in)

							leave := x.Source(in)
							defer leave()

							for i := 0; i <= b.N; i++ {
								in <- i
							}
						}()
					}

					done := make(chan struct{})
					go func(ctx context.Context) {
						defer close(done)
						ch, leave := x.Subscribe()
						defer leave()

						for {
							select {
							case <-ch:
							case <-ctx.Done():
								return
							}
						}
					}(ctx)

					wg.Wait()
					cancel()
					<-done
				})
			}
		})

		b.Run("outputs", func(b *testing.B) {
			for _, rr := range runners {
				rr(b, func(b *testing.B, n int) {
					x := New[int, int](Pass[int])

					ctx := context.Background()
					ctx, cancel := context.WithCancel(ctx)

					wg := sync.WaitGroup{}

					for i := 0; i < n; i++ {
						wg.Add(1)
						go func(ctx context.Context) {
							defer wg.Done()

							ch, leave := x.Subscribe()
							defer leave()

							for {
								select {
								case <-ch:
								case <-ctx.Done():
									return
								}
							}
						}(ctx)
					}

					in := make(chan int)
					defer close(in)

					leave := x.Source(in)
					defer leave()

					for i := 0; i < b.N; i++ {
						in <- i
					}

					cancel()
					wg.Wait()
				})
			}
		})

		b.Run("input-output", func(b *testing.B) {
			for _, runner := range runners {
				runner(b, func(b *testing.B, n int) {
					x := New[int, int](Pass[int])

					ctx := context.Background()
					ctx, cancel := context.WithCancel(ctx)

					og := sync.WaitGroup{}

					for i := 0; i < n; i++ {
						og.Add(1)
						go func(ctx context.Context) {
							defer og.Done()

							ch, leave := x.Subscribe()
							defer leave()

							for {
								select {
								case <-ch:
								case <-ctx.Done():
									return
								}
							}
						}(ctx)
					}

					sg := sync.WaitGroup{}

					for i := 0; i <= n; i++ {
						sg.Add(1)
						go func() {
							defer sg.Done()

							in := make(chan int)
							defer close(in)

							leave := x.Source(in)
							defer leave()

							for i := 0; i <= b.N; i++ {
								in <- i
							}
						}()
					}

					sg.Wait()
					cancel()
					og.Wait()
				})
			}
		})
	})

	b.Run("dynamic", func(b *testing.B) {
		b.Skip("implement me")
	})
}
