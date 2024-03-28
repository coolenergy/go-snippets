package mem

import "fmt"

// This shows if we take a set of a slice, are we taking the reference, or a copy?
func SubSliceSnippet() {
	a := []int{1, 2, 3, 4, 5, 5}
	b := a[2:5]

	fmt.Printf("%s\n", a)
	b[0] = 123
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", a)

}
