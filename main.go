package main

import(
  "fmt"
  c "action-test/combinatorics"
)

func main() {
  fmt.Println("Hello World!")
  var a1 = []int{0, 1, 2, 3}
	p1 := c.Perm2(a1)
	fmt.Printf("Permutation1: \n%v ", p1)
	fmt.Printf("Nr. of permutations: %v\n", len(p1))
}
