package contact

import (
	"testing"
)

func benchmarkjoinstring(b *testing.B) {
	str := []string{"hello", ",", "string", "!"}

	for i := 0; i < b.N; i++ {
		joinstrings(str)
	}
}
