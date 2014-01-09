package ballclock

import "testing"

func TestIt(t *testing.T) {
	c := New(30)
	p := c.Period()
	if p != 15 {
		t.Fatalf("too bad %d not 15", p)
	}
}
