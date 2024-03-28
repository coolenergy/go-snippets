package mem

import "fmt"

// This snippet shows that a slice is created if we use the [] operators
// so it is not a reference, but a whole new sub copy
func stringSliceAddress() {
	first := "First"
	firstSlice := first[0:3]

	fmt.Printf("First address: %p\nSlice address: %p\n",
		&first, &firstSlice)

	second := new(string)
	*second = "Second String"
	secondSlice := (*second)[0:3]

	fmt.Printf("Second address: %p\nSecond address: %p\n",
		second, &secondSlice)
}
