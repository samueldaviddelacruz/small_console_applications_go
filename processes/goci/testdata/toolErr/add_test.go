package add

import "testing"

func TestAdd(t *testing.T) {
	a := 2
	b := 3
	expected := 5
	res := add(a, b)
	if res != expected {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
