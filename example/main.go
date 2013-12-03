package main

import "fmt"
import "github.com/shanemhansen/ballclock"
import "sync"
import "os"
import "bufio"
import "io"
import "strconv"

func main() {
    w := bufio.NewReader(os.Stdin)
    data := make([]int, 0, 4)
    for {
        line, err := w.ReadString('\n')
        if err == io.EOF {
            break
        }
        line = line[:len(line)-1]
        number, err := strconv.ParseInt(line, 10, 0)
        if err != nil {
            panic("Invalid number"+line)
        }
        data = append(data, int(number))
    }
    results := make(chan [2]int)
    var wg sync.WaitGroup
    for _, j := range data {
        wg.Add(1)
        go func(period int) {
            defer wg.Done()
            template:= ballclock.BallClock{Polynomial: []int{5-1,60/5-1,12-1}, QueueSize: period}
            clock, err := ballclock.New(&template)
            if err != nil {
                //per spec, skip                
                return
            }
            i := clock.Period()
            results <- [2]int{period, i}
        }(j)
    }
    go func() {
        for {
            result := <- results
            if result[0] == 0 && result[1] == 0 {
                return
            }
            fmt.Printf("%d balls cycle after %d Days\n", result[0], (result[1]+1)/60/24)
        }
    }()
    wg.Wait()
    results <- [2]int{0,0}
    
}
