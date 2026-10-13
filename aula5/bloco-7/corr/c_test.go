package corr

import "testing"

func TestOk(t *testing.T) {
	n := 0
	n++
	if n != 1 {
		t.Fatal(n)
	}
}
