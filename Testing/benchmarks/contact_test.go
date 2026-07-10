package contact

import (
	"testing"
)

func BenchmarkJoinstring(b *testing.B) {
	str := []string{"hello", ",", "string", "!"}

	for i := 0; i < b.N; i++ {
		Joinstrings(str)
	}
}
