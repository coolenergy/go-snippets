package mem

import "fmt"

// This snippet shows that an interface{} maybe a literal or a pointer
// It won't do some black magic casting or copying
func printInterfacePtrAndValue(some interface{}) {
	fmt.Printf("[PrintString] %p %s\n", some, some)
}

func StringCopyInterface() {
	first := "First"
	fmt.Printf("[Main] %p %s\n", first, first)
	printInterfacePtrAndValue(first)
	fmt.Println()

	second := "Second"
	fmt.Printf("[Main] %p %s\n", &second, &second)
	printInterfacePtrAndValue(&second)
	fmt.Println()

	third := new(string)
	*third = "Third"
	fmt.Printf("[Main] %p %s\n", third, third)
	printInterfacePtrAndValue(third)

}
