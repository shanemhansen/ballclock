package ballclock

import "testing"

type Strategy func(*BallClock) int
// Test that we get proper result as a function of queue size
// With known data, and all algorithms
func TestPeriod(t *testing.T) {
	tests := [][2]int{
		{30, 15},
		{45, 378},
	}
	for _, testcase := range tests {
		expected := testcase[1]
        algorithms := []Strategy{(*BallClock).Period, (*BallClock).PeriodBySuccessivePermuations}
        for _, algo := range algorithms {
	        clock, err :=  New(&BallClock{Polynomial: []int{5-1, 60/5-1, 12-1}, QueueSize: testcase[0]})
            if err != nil {
                t.Fatal(err)
            }
            if ! clock.IsPristine() {
                t.Fatal("precondition failed")
            }
            actual := algo(clock)
            actual/=(60*24)
            if expected != actual {
                t.Fatalf("Expected %d got %d", expected, actual)
            }
        }
	}
}
