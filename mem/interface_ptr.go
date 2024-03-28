package mem

import "fmt"

//(1) The interface
type Mutable interface {
	mutate(newValue string) error
}

//(2) Struct
type Data struct {
	name string
}

//(3) Implements the interface with a pointer receiver
func (d *Data) mutate(newValue string) error {
	d.name = newValue
	return nil
}

//(4) Function that accepts the interface
func mutator(mute Mutable) error {
	fmt.Printf("mutator: Address is %p\n", mute)
	return mute.mutate("mutate")
}

func InterfacePtr() {
	d := Data{name: "fresh"}
	fmt.Printf("mutator: Address is %p\n", &d)
	fmt.Println(d.name) //fresh
	//(5) pass as a pointer
	mutator(&d)
	fmt.Println(d.name) //mutate
}
