package bench

import (
	"context"
	"testing"
	"time"
)

var (
	ctx        = context.Background()
	ctxWithVal = context.WithValue(context.Background(), 1, 2)
)

func BenchmarkCtx_Select(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-ctx.Done():
		default:
		}
	}
}

func BenchmarkCtx_Error(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctx.Err()
	}
}

func BenchmarkCtx_SelectVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		select {
		case <-ctxWithVal.Done():
		default:
		}
	}
}

func BenchmarkCtx_ErrorVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ctxWithVal.Err()
	}
}

func BenchmarkCtx_SelectWithCanceled(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case <-ctx.Done():
		default:
		}
	}
}

func BenchmarkCtx_ErrorWithCacneled(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Err()
	}
}

func BenchmarkCtx_SelectWithTimeout(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		select {
		case <-ctx.Done():
		default:
		}
	}
}

func BenchmarkCtx_ErrorWithTimeout(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx.Err()
	}
}
