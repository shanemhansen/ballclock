package ballclock

import "fmt"

var _ fmt.Stringer

type Clock struct {
	queuesize int
	queue     []int
	arm1pos   int
	arm1      [4]int
	arm2pos   int
	arm2      [11]int
	arm3pos   int
	arm3      [11]int
}

func New(queuesize int) *Clock {
	q := make([]int, queuesize)
	for i := 0; i < queuesize; i++ {
		q[i] = i + 1
	}
	return &Clock{
		queuesize: queuesize,
		queue:     q,
	}
}
func (c *Clock) Period() int {
	i := 0
TICK:
	for {
		i++
		ball := c.queue[0]
		c.queue = c.queue[1:]
		// 1 minute
		if c.arm1pos != len(c.arm1) {
			c.arm1[c.arm1pos] = ball
			c.arm1pos++
			continue
		}
		c.arm1pos = 0
		c.queue = append(
			c.queue, c.arm1[3], c.arm1[2], c.arm1[1], c.arm1[0],
		)
		// 5 minute
		if c.arm2pos != len(c.arm2) {
			c.arm2[c.arm2pos] = ball
			c.arm2pos++
			continue
		}
		c.arm2pos = 0
		c.queue = append(
			c.queue,
			c.arm2[10], c.arm2[9], c.arm2[8],
			c.arm2[7], c.arm2[6], c.arm2[5], c.arm2[4],
			c.arm2[3], c.arm2[2], c.arm2[1], c.arm2[0],
		)
		// 1 hour
		if c.arm3pos != len(c.arm3) {
			c.arm3[c.arm3pos] = ball
			c.arm3pos++
			continue
		}
		c.arm3pos = 0
		c.queue = append(
			c.queue,
			c.arm3[10], c.arm3[9], c.arm3[8],
			c.arm3[7], c.arm3[6], c.arm3[5], c.arm3[4],
			c.arm3[3], c.arm3[2], c.arm3[1], c.arm3[0],
			ball,
		)
		if len(c.queue) == c.queuesize {
			// Could be...
			for i, k := range c.queue {
				if i+1 != k {
					continue TICK
				}
			}
			return i / (24 * 60)
		}

	}
}
