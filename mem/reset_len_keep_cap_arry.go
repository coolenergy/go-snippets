package mem

import "fmt"

// This snippets shows how to reuse the very same array by recycling it
func ResetLenKeepCapArray() {
	example := []string{"A", "B", "C", "D", "E"}
	fmt.Println(fmt.Sprintf("%p", example), example, len(example), cap(example))
	example[0] = "1"
	example[2] = "3"
	fmt.Println(fmt.Sprintf("%p", example), example, len(example), cap(example))
	example = example[:0]
	fmt.Println(fmt.Sprintf("%p", example), example, len(example), cap(example))
	example = append(example, "1")
	fmt.Println(fmt.Sprintf("%p", example), example, len(example), cap(example))
	example = append(example, "2")
	fmt.Println(fmt.Sprintf("%p", example), example, len(example), cap(example))
}
