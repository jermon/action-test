package actiontest

//import "fmt"

///////////////////////////////////////////////////////////////////////////////

var permList [][]int

// Perm2 permutes a slice if ints.
func Perm2(a []int) [][]int {
	perm2(a, 0)
	return permList
}

// Permute the values at index i to len(a)-1.
func perm2(a []int, i int) {
	if i > len(a) {
		// fmt.Printgithub.com/jermon/action-testln(a)
		b := make([]int, len(a))
		copy(b, a)
		permList = append(permList, b)
		return
	}
	perm2(a, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm2(a, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Perm3 r from n permutation
// No repititions.
func Perm3(a []int, i int) [][]int {
	// fmt.Printf("** I: %v, A: %v\n", i, a)
	var result [][]int

	for i1, v1 := range a {

		if i > 1 {
			var reminder []int
			reminder = append(reminder, a[:i1]...)
			reminder = append(reminder, a[i1+1:]...)
			// fmt.Printf("V: %v, Rem: %v\n", v1, reminder)

			part := Perm3(reminder, i-1)

			for _, v2 := range part {
				line := make([]int, 0)
				line = append(line, v1)
				// fmt.Printf("## %v, %v\n", line, v2)
				line = append(line, v2...)
				result = append(result, line)
			}

		} else {
			line := make([]int, 0)
			line = append(line, v1)
			// fmt.Printf("Ading line: %v\n",line)
			result = append(result, line)
		}
	}
	// fmt.Printf("Result: %v\n",result)
	return result
}

// NrOfPermutations returns the number of permutations r from n.
func NrOfPermutations(r, n int) int {
	return factorial(n) / factorial(n-r)
}
