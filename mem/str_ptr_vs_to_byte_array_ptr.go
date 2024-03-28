package mem

import "fmt"

/*
This snippet shows that []byte() creates a copy of a string
*/
func StrPtrVsToByteArrayPtr() {

	dummyText := "Some dummy text"

	fmt.Printf("Address of string %p.\nAddress Of []byte %p",
		&dummyText, []byte(dummyText))

}
