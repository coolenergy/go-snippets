package mem

import "fmt"

// This snippets shows that strings are passed by value
// But that does not mean the string is copied, what is copied is the ptr to the string
// and not the string itself, so we are safe using func a(b string) there is no performance penalty
// life if we were in C++ where the string would be copied completely
func printString(myStringPtr string) {
	fmt.Printf("%p\n", &myStringPtr)
}

func printStringPtr(myStringPtr *string) {
	fmt.Printf("%p\n", myStringPtr)
}

func StringByRef() {
	firstStr := "First string"
	fmt.Printf("%p\n", &firstStr)
	printStringPtr(&firstStr)

	secondStr := new(string)
	*secondStr = "Second string"
	fmt.Printf("%p\n", secondStr)
	printStringPtr(secondStr)

}
