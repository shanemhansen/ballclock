package ballclock

import "testing"

// Test that we get proper result as a function of queue size
func TestPeriod(t *testing.T) {
	tests := [][2]int{
		{30, 15},
		{45, 378},
	}
	for _, testcase := range tests {
		expected := testcase[1]
	    clock, err := New(&BallClock{Polynomial: []int{5-1, 60/5-1, 12-1}, QueueSize: testcase[0]})
        if err != nil {
            t.Fatal(err)
        }
        actual := clock.Period()
        actual++
        actual/=(60*24)
        if expected != actual {
            t.Fatalf("Expected %d got %d", expected, actual)
        }
	}
}
