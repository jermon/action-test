package combinatorics

import (
	"fmt"
	"testing"
)

func TestNrOfCombinations(t *testing.T) {

	tests := []struct {
		n      int
		r      int
		output int
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 1, 3},
		{4, 1, 4},
		{2, 2, 1},
		{3, 2, 3},
		{4, 2, 6},
		{3, 3, 1},
		{4, 3, 4},
		{4, 4, 1},
		{5, 2, 10},
		{6, 3, 20},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test-%d", i), func(t *testing.T) {
			if got := NrOfCombinations(tt.r, tt.n); got != tt.output {
				t.Errorf("NrOfCombinations(%v, %v) = %d; want %v", tt.r, tt.n, got, tt.output)
			}
		})
	}
}
