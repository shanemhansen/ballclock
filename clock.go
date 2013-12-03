// A ballclock represents a set of 3 rungs
// The first rung represents minutes
// The second 5 minute increments
// and the last, 1 hour increments
// In the general case, the state of a ball clock
// can be defined as a polynomial using a variable base
// (10/1, 60/5, 12/1) or (10, 2, 12)
// We can defined a function t(x1, x2, x3) where
// t is time in minutes, x1 is minutes, x2 is sets of minutes
// and x3 is hours. as t(x1, x2, x3) = x1 + 5*x2 + 60*x3
// where  x1 ∈ [0..9], x2 ∈ [0,12], x3 ∈ [0, 12]
package ballclock

import "errors"
var InvalidClock error = errors.New("Invalid clock, not enough balls!")
//General polynomial type
type Polynomial []int

// Ballclock is a specific type of polynomial
type BallClock struct {
    Polynomial
    QueueSize int
    queue []int
    arms [][]int
    pristineQueue []int
}
func New(template *BallClock) (*BallClock, error) {
    template.queue = make([]int, template.QueueSize)
    template.pristineQueue = make([]int, template.QueueSize)
    for i:=0; i< template.QueueSize; i++ {
        template.queue[i] = i+1
    }
    copy(template.pristineQueue, template.queue)
    required := 0
    for _, size := range template.Polynomial {
        required += size
        template.arms = append(template.arms, make([]int, size))
    }
    if template.QueueSize < required {
        return nil, InvalidClock
    }
    return template, nil
}

// Determine the Period of a wall clock before it repeats itself
// in the order
func (self *BallClock) Period() int {
    var i int
    self.Tick()
    for i=0; !self.IsPristine(); i++ {
        self.Tick()
    }
    return i
}

func (self *BallClock) State() ([]int, [][]int) {
    return self.queue, self.arms
}
func (self *BallClock) IsPristine() bool {
    if len(self.pristineQueue) != len(self.queue) {
        return false
    }
    for i, pristine := range self.pristineQueue {
        if self.queue[i] != pristine {
            return false
        }
    }
    return true
}
func (self *BallClock) Tick() {
    // Simple counting algorithm, with a twist
    ball := self.queue[0]
    self.queue = self.queue[1:]
    armnumber :=0
    for {
        if armnumber == len(self.arms) {
            //Put the ball back in the queue, at the end
            self.queue = append(self.queue, ball)
            return
        }
        arm := self.arms[armnumber]
        //find a place to put the ball
        for i, val := range arm {
            if val == 0 {
                arm[i] = ball
                return
            }
        }
        // uhoh, nowhere to put the ball
        // time to carry, loop over arm backwards
        // empyting it out and placing the results
        // in the queue
        for i, _ := range arm {
            carry := &arm[len(arm)-i-1]
            if *carry == 0 {
                panic("wtf")
            }
            self.queue = append(self.queue, *carry)
            *carry = 0
        }
        //let's try the next arm.. unless we're out of arms?
        armnumber++            
    }
    panic("uh oh")
}
