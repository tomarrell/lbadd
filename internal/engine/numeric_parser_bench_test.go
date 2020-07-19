package engine

import (
	"testing"

	"github.com/tomarrell/lbadd/internal/engine/types"
)

func BenchmarkToNumericValue(b *testing.B) {
	str := "75610342.92389E-21423"
	expVal := 75610342.92389E-21423

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		val, ok := ToNumericValue(str)
		if !ok {
			b.FailNow()
		}
		if expVal != val.(types.RealValue).Value {
			b.FailNow()
		}
	}
}
