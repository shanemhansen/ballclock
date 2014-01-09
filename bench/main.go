package main

import "fmt"
import "github.com/shanemhansen/ballclock"
import "github.com/davecheney/profile"

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	balls := 123
	c := ballclock.New(balls)
	for i := 1; i < 10; i++ {
		fmt.Printf("%d balls cycle in %d days\n", balls, c.Period())
	}
}
