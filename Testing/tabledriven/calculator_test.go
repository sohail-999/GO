package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{5, 6, 11},
		{10, 20, 30},
		{6, 7, 13},
	}

	for _, tt := range tests {
		got := Add(tt.a, tt.b)

		if got != tt.want {
			t.Errorf("Add(%d,%d)=%d want %d",
				tt.a,
				tt.b,
				got,
				tt.want,
			)
		}
	}
}
