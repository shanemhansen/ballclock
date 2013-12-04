package main

import "io"
import "os"
import "fmt"
import "flag"
import "sync"
import "bufio"
import "strconv"

import "github.com/shanemhansen/ballclock"

var algo string

// read integers from stdin, returning text describing the period.
// output may arrive out of order, yay goroutines.
func main() {
	flag.StringVar(&algo, "algo", "iterateminutes", "algorithm iterateminutes|permutedays|lcm")
	strategies := map[string]func(*ballclock.BallClock) int{
		"lcm":            (*ballclock.BallClock).PeriodByLCM,
		"iterateminutes": (*ballclock.BallClock).Period,
		"permutedays":    (*ballclock.BallClock).PeriodBySuccessivePermuations}
	flag.Parse()
	strategy := strategies[algo]
	if strategy == nil {
		panic(algo + " is not a valid strategy")
	}
	w := bufio.NewReader(os.Stdin)
	data := make([]int, 0, 4)

	// Read lines
	for {
		line, err := w.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = line[:len(line)-1]
		number, err := strconv.ParseInt(line, 10, 0)
		if err != nil {
			// Be a little robust, skip empty spaces and whatnot
			continue
		}
		data = append(data, int(number))
	}
	// Easy way to fire off lots of goroutines and wait for them to finish
	// Rob Pike told me this is the equivalant of & for unix shells.
	results := make(chan [2]int)
	var wg sync.WaitGroup
	for _, j := range data {
		wg.Add(1)
		go func(period int) {
			defer wg.Done()
			template := ballclock.BallClock{Polynomial: []int{5 - 1, 60/5 - 1, 12 - 1}, QueueSize: period}
			clock, err := ballclock.New(&template)
			if err != nil {
				//per spec, skip
				return
			}
			i := strategy(clock)
			results <- [2]int{period, i}
		}(j)
	}
	// Asynchronously collect results while waiting for the process to finish
	go func() {
		for {
			result := <-results
			if result[0] == 0 && result[1] == 0 {
				return
			}
			fmt.Printf("%d balls cycle after %d Days\n", result[0], (result[1]))
		}
	}()
	//Join worker goroutines
	wg.Wait()
	// Done, send a quit message to the reader
	results <- [2]int{0, 0}
}
