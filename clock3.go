package ballclock

// The solution to the problem now looks like:
// 1. Model the clock
// 2. Move the dial forward 24 hours (or 12 if you prefer)
// 3. Examine the clock state, the queue is a permutation
//     vector for 24 hours.
// 4. Write the permutation in factorized form
// 5. Compute the LCM of the elements of the factorized form
// 6. Return that as the number of days in the period
// Algorithmically this is somewhat tricky. The clock operation
// is actually constant time w/ respect to the size of the queue.
// it merely requires a fixed set of iterations. O(1). Now
// with some smarter permutation work involving powers of
// the permutation we could decrease that constant, but I'll
// leave that out. This is of course the famouse russion method.
// Factorizing the permutation using my naive algorithm is
// probably O(n**2) as a function of queue size.
// I loop over the permutations, requiring a maximum of one loop
// per cycle, and there can be up to n cycles in the identity permutation
// Caclulating the LCM is extremely efficient. It's been shown
// that euclids famous algorithm runs in logarithmatic time,
// So that adds O(n*lg n) steps.
// Enough talking
func (self *BallClock) PeriodByLCM() int {
	vec := self.GetPermutationVector(DAY)
	period := Cycle(vec)
	return period
}
