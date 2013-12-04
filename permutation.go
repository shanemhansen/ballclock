package ballclock

// Permute a vector by the given permutation vector
// Leaving the original unchanged
func Permute(permutation, input []int) (output []int) {
	if len(input) != len(permutation) {
		panic("Permutation and input lengths don't match")
	}
	output = make([]int, len(input))
	for i, k := range permutation {
		output[i] = input[k]
	}
	return output
}

// Return a vector factored into cyclic form
// (a slice of slices of ints) The basic strategy
// is to pick a number, follow it until it leads back,
// the whole time burning the bridge behind us (replacing with -1)
// When we've destroyed the permutation vector (all -1's), we're done.
func CyclicForm(p []int) (cycles [][]int) {
	cycles = make([][]int, 0)
	permutations := make([]int, len(p))
	copy(permutations, p)
	head := permutations[0]
	cycle := []int{head}
	next := head
	marker := -1
THE_CYCLE:
	for {
		// find a single cycle, stashing a ref
		// so that we can clean up after we've
		// got our next target
		oldnext := next
		next = permutations[next]
		permutations[oldnext] = marker
		if next == head {
			// The cycle is complete. Find
			// a new starting point if possible
			for _, k := range permutations {
				if k == marker {
					continue
				}
				next = k
				head = k
				cycles = append(cycles, cycle)
				cycle = []int{head}
				continue THE_CYCLE
			}
			// Otherwise let's call it a day
			cycles = append(cycles, cycle)
			break THE_CYCLE
		}
		cycle = append(cycle, next)
	}
	return cycles
}

// return the number of iterations until the input cycles
// This is known, and by intuitively seems to be the
// least common multime of the length of all the involved cycles
func Cycle(permutation []int) int {
	cycles := CyclicForm(permutation)
	acc := 1
	for _, cycle := range cycles {
		acc = lcm(acc, len(cycle))
	}
	return acc
}

// Good old euclidian algorithm for gcd
// Source: Euclid's elements Books VII and X
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// and it's closely related cousin least common multiple
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
