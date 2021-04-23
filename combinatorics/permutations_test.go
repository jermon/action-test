package combinatorics

import (
	"fmt"
	"testing"
)

func TestPerm2(t *testing.T) {
	var a1 = []int{0, 1, 2, 3}
	p1 := Perm2(a1)
	t.Logf("Permutation1: \n%v ", p1)
	t.Logf("Nr. of permutations: %v\n", len(p1))

}

func TestPerm(t *testing.T) {
	var a1 = []int{0, 1, 2, 3}
	p1 := Perm3(a1, 2)
	t.Logf("Permutation1: \n%v ", p1)
	t.Logf("Nr. of permutations: %v\n", len(p1))

}

func TestNrOfPermutations(t *testing.T) {

	tests := []struct {
		n      int
		r      int
		output int
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 1, 3},
		{4, 1, 4},
		{2, 2, 2},
		{3, 2, 6},
		{4, 2, 12},
		{3, 3, 6},
		{4, 3, 24},
		{4, 4, 24},
		{5, 2, 20},
		{6, 3, 120},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test-%d", i), func(t *testing.T) {
			if got := NrOfPermutations(tt.r, tt.n); got != tt.output {
				t.Errorf("NrOfPermutations(%v, %v) = %d; want %v", tt.r, tt.n, got, tt.output)
			}
		})
	}
}
