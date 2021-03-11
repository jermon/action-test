package combinatorics

import (
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

func TestFactorial0(t *testing.T) {
    input := 0
    expected := 1
    got := Factorial(input)
    if got != expected {
        t.Errorf("Factorial(%v) = %d; want %v",input, got, expected)
    }
}

func TestFactorial1(t *testing.T) {
    input := 1
    expected := 1
    got := Factorial(input)
    if got != expected {
        t.Errorf("Factorial(%v) = %d; want %v",input, got, expected)
    }
}

func TestFactorial2(t *testing.T) {
    input := 2
    expected := 2
    got := Factorial(input)
    if got != expected {
        t.Errorf("Factorial(%v) = %d; want %v",input, got, expected)
    }
}

func TestFactorial3(t *testing.T) {
    input := 3
    expected := 6
    got := Factorial(input)
    if got != expected {
        t.Errorf("Factorial(%v) = %d; want %v",input, got, expected)
    }
}

func BenchmarkFactorial(b *testing.B) {
    input := 3
      for i := 0; i < b.N; i++ {
        Factorial(input)
      }
}

func TestNrOfPermutations1(t *testing.T) {
  in1, in2 := 1,2
  expected := 2
  got := NrOfPermutations(in1,in2)
  if got != expected {
    t.Errorf("NrOfPermutations(%v,%v) = %d; want %v",in1,in2, got, expected)
  }
}

func TestNrOfPermutations2(t *testing.T) {
  in1, in2 := 2, 5
  expected := 20
  got := NrOfPermutations(in1,in2)
  if got != expected {
    t.Errorf("NrOfPermutations(%v,%v) = %d; want %v",in1,in2, got, expected)
  }
}

func TestNrOfPermutations3(t *testing.T) {
  in1, in2 := 3, 6
  expected := 120
  got := NrOfPermutations(in1,in2)
  if got != expected {
    t.Errorf("NrOfPermutations(%v,%v) = %d; want %v",in1,in2, got, expected)
  }
}



