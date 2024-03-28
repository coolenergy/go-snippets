package mem

import (
	"fmt"
	"time"
)

// This shows if a function returns a slice, is it getting copied
// or a reference is returned?
func getSlice() []int {
	a := []int{1, 2, 3, 4, 5, 5}
	go func() {
		time.Sleep(time.Second)
		fmt.Printf("%p %s\n", a, a)
	}()
	fmt.Printf("%p %s\n", a, a)
	return a[2:5]
}

func SubSliceReturnSnippet() {
	a := getSlice()
	a[0] = 123
	fmt.Printf("%p %s\n", a, a) // the address is addressOf(original) + sizeOf(int)
	time.Sleep(time.Second * 3)
}
