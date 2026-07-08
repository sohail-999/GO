package calculator

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 2, -1},
		{5, 6, -1},
		{10, 20, -10},
	}

	for _, tt := range tests {
		got := Subtract(tt.a, tt.b)

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
