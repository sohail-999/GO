package calculator

import "testing"

func TestAdd(t *testing.T) {
	x, y := 5, 3
	expected := 8
	result := Add(x, y)
	if result != expected {
		t.Errorf("add(%d ,%d)=%d,expected%d", x, y, result, expected)

	}

}
func TestSubtract(t *testing.T) {
	x, y := 5, 3
	expected := 2
	result := Subtract(x, y)
	if result != expected {
		t.Errorf("add(%d ,%d)=%d,expected%d", x, y, result, expected)

	}

}
