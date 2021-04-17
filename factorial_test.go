package combinatorics

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {

	tests := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test-%d", i), func(t *testing.T) {
			if got := factorial(tt.input); got != tt.output {
				t.Errorf("factorial(%v) = %d; want %v", tt.input, got, tt.output)
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	input := 99
	for i := 0; i < b.N; i++ {
		factorial(input)
	}
}
