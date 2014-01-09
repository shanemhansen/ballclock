package main

import "fmt"
import "github.com/shanemhansen/ballclock"

func main() {
	balls := 123
	c := ballclock.New(balls)
	fmt.Printf("%d balls cycle in %d days\n", balls, c.Period())
}
