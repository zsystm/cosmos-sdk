package context

import (
	"context"
	"testing"
)

func BenchmarkContext(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		ctx = context.WithValue(ctx, i, i)
	}

	y := 0
	for i := 0; i < b.N; i++ {
		x := ctx.Value(0).(int)
		y += x
	}
	b.Logf("%d", y)
}

func BenchmarkMap(b *testing.B) {
	m := map[int]int{}
	for i := 0; i < 100; i++ {
		m[i] = i
	}

	y := 0
	for i := 0; i < b.N; i++ {
		x := m[0]
		y += x
	}
	b.Logf("%d", y)
}
