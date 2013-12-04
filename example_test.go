package ballclock_test

import "github.com/shanemhansen/ballclock"
import "fmt"

func ExampleNew() {
	template := ballclock.BallClock{Polynomial: []int{5 - 1, 60/5 - 1, 12 - 1}, QueueSize: 30}
	clock, err := ballclock.New(&template)
	if err != nil {
		panic(err)
	}
	fmt.Println(clock.Period())
	// Output:
	// 15
}

func ExampleBallClock_GetPermutationVector() {
	template := ballclock.BallClock{Polynomial: []int{5 - 1, 60/5 - 1, 12 - 1}, QueueSize: 30}
	clock, err := ballclock.New(&template)
	if err != nil {
		panic(err)
	}
	ticks := ballclock.DAY // 1 day
	vec := clock.GetPermutationVector(ticks)
	fmt.Println(len(vec))
	// Output:
	// 30
}
func ExamplePermute() {
	leftshift := []int{1, 2, 3, 4, 0}
	input := []int{3, 1, 4, 5, 6}
	output := ballclock.Permute(leftshift, input)
	fmt.Println(output)
	// Output:
	// [1 4 5 6 3]
}

func ExampleCyclicForm() {
	swap := []int{1, 0, 3, 4, 2}
	fmt.Println(ballclock.CyclicForm(swap))
	// Output:
	// [[1 0] [3 4 2]]
}

func ExampleCycle() {
	swap := []int{1, 0, 3, 4, 2}
	fmt.Println(ballclock.Cycle(swap))
	// Output:
	// 6
}
