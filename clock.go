package ballclock

import "fmt"

var _ fmt.Stringer

type Clock struct {
	queuesize int
	queue     *Queue
	arm1pos   int
	arm1      [4]int
	arm2pos   int
	arm2      [11]int
	arm3pos   int
	arm3      [11]int
}
type Queue struct {
	queue []int
	head  int
	tail  int
}

func (q *Queue) Set(i, val int) {
	if q.head != 0 {
		panic("no no no no no, you can't set values after the fact")
	}
	q.queue[i] = val
}
func (q *Queue) Pop() (result int) {
	result = q.queue[q.head]
	q.head++
	l := len(q.queue)
	if q.head >= l {
		q.head -= l
	}
	return
}
func (q *Queue) Len() int {
	t := q.tail - q.head
	if t > 0 {
		return t
	}
	return (len(q.queue) - t)
}
func (q *Queue) append(vals ...int) {
	// TODO implement contig
	l := len(q.queue)
	vlen := len(vals)
	if l-q.tail > vlen {
		copy(q.queue[q.tail:], vals)
		q.tail += vlen
		return
	}
	v := 0
	// from middle to end
	for i := q.head; i < l && v < vlen; i++ {
		q.queue[q.tail] = vals[v]
		v++
		q.tail++
		if q.tail >= l {
			q.tail -= l
		}
	}
	if v == vlen {
		return
	}
	// Copy from begining to end
	// No check needed for buffer overflow
	for i := 0; v < vlen; i++ {
		q.queue[q.tail] = vals[v]
		v++
		q.tail++
		if q.tail >= l {
			q.tail -= l
		}
	}
	//	q.queue = append(q.queue, vals...)
}
func NewQueue(queuesize int) *Queue {
	return &Queue{queue: make([]int, queuesize)}
}
func New(queuesize int) *Clock {
	q := NewQueue(queuesize)
	for i := 0; i < queuesize; i++ {
		q.Set(i, i)
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
		ball := c.queue.Pop()
		// 1 minute
		if c.arm1pos != len(c.arm1) {
			c.arm1[c.arm1pos] = ball
			c.arm1pos++
			continue
		}
		c.arm1pos = 0
		c.queue.append(c.arm1[3], c.arm1[2], c.arm1[1], c.arm1[0])
		// 5 minute
		if c.arm2pos != len(c.arm2) {
			c.arm2[c.arm2pos] = ball
			c.arm2pos++
			continue
		}
		c.arm2pos = 0
		c.queue.append(
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
		c.queue.append(
			c.arm3[10], c.arm3[9], c.arm3[8],
			c.arm3[7], c.arm3[6], c.arm3[5], c.arm3[4],
			c.arm3[3], c.arm3[2], c.arm3[1], c.arm3[0],
			ball,
		)
		if i%1440 == 0 && c.queue.Len() == c.queuesize {
			// Could be...
			for i, k := range c.queue.queue {
				if i != k {
					continue TICK
				}
			}
			return i / (24 * 60)
		}
	}
}
