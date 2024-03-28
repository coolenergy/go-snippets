package mem

import (
	"fmt"
	"unsafe"
)

func addFruit(fruits []string) unsafe.Pointer {
	fmt.Printf("Array address %p %d %d\n", fruits, len(fruits), cap(fruits))
	fruits = append(fruits, "Strawberry")
	fmt.Printf("Array address %p %d %d\n", fruits, len(fruits), cap(fruits))
	return unsafe.Pointer(&fruits[2])
}

func SliceParamReturn() {
	fruits := make([]string, 0, 5)
	fmt.Printf("Array address %p %d %d\n", fruits, len(fruits), cap(fruits))
	fruits = append(fruits, "Apple")
	fruits = append(fruits, "Banana")

	fmt.Printf("Array address %p %d %d\n", fruits, len(fruits), cap(fruits))
	ptr := addFruit(fruits)
	fmt.Printf("Array address %p %d %d\n", fruits, len(fruits), cap(fruits))
	fmt.Println(*(*string)(ptr))

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}
