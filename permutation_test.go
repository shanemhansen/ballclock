package ballclock

import "testing"

// test obtaining the 24 hour permutation vector
func TestPermutationVector(t *testing.T) {
	// Assert that the permutation vector for
	// a cycle is the identity permutation
	clock, err := New(&BallClock{Polynomial: []int{5 - 1, 60/5 - 1, 12 - 1}, QueueSize: 30})
	if err != nil {
		t.Fatal(err)
	}
	var cycle Day = DAY * 15
	vec := clock.GetPermutationVector(cycle)
	for i, k := range vec {
		if i != k {
			t.Fatalf("Expected identity vector element %d got %d", i, k)
		}
	}
}

// test we know how to apply a permutation
func TestPermute(t *testing.T) {
	permutations := [][2][]int{
		{{1, 2, 0}, {5, 6, 4}},
		{{0, 1, 2}, {4, 5, 6}}}
	input := []int{4, 5, 6}
	for _, testcase := range permutations {
		permutation := testcase[0]
		expected := testcase[1]
		output := Permute(permutation, input)
		for i, k := range expected {
			if k != output[i] {
				t.Fatalf("Permutation function failed, expected %d got %d", k, output[i])
			}
		}
	}
}

// Ensure we can properly put a permutation into cyclic form
// We make no guarantees about the ordering within the groups
func TestGetCyclicForm(t *testing.T) {
	//from el wikipedia
	input := []int{3, 1, 6, 5, 4, 7, 0, 2}
	expected := [][]int{{0, 3, 5, 7, 2, 6}, {1}, {4}}

	actual := CyclicForm(input)
	if len(actual) != len(expected) {
		t.Fatalf("decomposition failed")
	}
	for i, got := range actual {
		if len(got) != len(expected[i]) {
			t.Fatalf("decomposition failed")
		}
	}
}

// Test we can calculate the cycle using the LCM approach
func TestGetLCD(t *testing.T) {
	if Cycle([]int{1, 0}) != 2 {
		t.Fatalf("expected %d, got %d", 4, Cycle([]int{2, 4}))
	}
	if Cycle([]int{1, 2, 0, 4, 3}) != 6 {
		t.Fatalf("expected %d, got %d", 4, Cycle([]int{2, 4}))
	}
}
