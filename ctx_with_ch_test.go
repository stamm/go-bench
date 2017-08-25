package bench

import (
	"context"
	"testing"
	"time"
)

var (
	timer     = 400 * time.Millisecond
	sleepTime = 500 * time.Millisecond
)

func BenchmarkCtxWithCh_Select(b *testing.B) {
	timer := time.NewTimer(timer)
	defer timer.Stop()

	ctx, cancel := context.WithCancel(context.Background())

	res := make(chan struct{})

	go func() {
		time.Sleep(sleepTime)
		select {
		case <-ctx.Done():
		case res <- struct{}{}:
		}
		close(res)
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case <-res:
		case <-timer.C:
			cancel()
		}
	}
}

func BenchmarkCtxWithCh_BufChan(b *testing.B) {
	timer := time.NewTimer(timer)
	defer timer.Stop()

	res := make(chan struct{}, 1)

	go func() {
		time.Sleep(sleepTime)
		res <- struct{}{}
		close(res)
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case <-res:
		case <-timer.C:
		}
	}
}
