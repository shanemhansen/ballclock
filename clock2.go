package ballclock

// clock2 holds implementtion of the permuatedays algorithm
type Day int

var DAY Day = 24 * 60

// Get the period, measured in days (24 hour permutations applied)
func (self *BallClock) PeriodBySuccessivePermuations() int {
	vec := self.GetPermutationVector(DAY)
	i := 1
	self.queue = Permute(vec, self.queue)
	for !self.IsPristine() {
		i++
		self.queue = Permute(vec, self.queue)
	}
	return i
}

// Return a permuted vector representing the
// changes made to the initial state. by ticks Resets BallClock
func (self *BallClock) GetPermutationVector(needed_ticks Day) []int {
	for i := Day(0); i < needed_ticks; i++ {
		self.Tick()
	}
	//get permuation vector
	vec, _ := self.State()
	if len(vec) != self.QueueSize {
		panic("Number of ticks does not form a permutation")
	}
	// Since we numbered the balls starting at 1 in order to leave room for a
	// sentinal 0 value, we need to adjust our permutation vector
	for i, _ := range vec {
		vec[i]--
	}
	New(self)
	return vec
}
