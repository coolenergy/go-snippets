package core

import "fmt"

func PanicRecover(x int) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Exception Handled in goroutine, error: %v\n", err)
		}
	}()

	fmt.Printf("Trying 2/%d\n", x)
	//noinspection GoDivisionByZero
	fmt.Println(2 / x)
	fmt.Println("After division ...")
}
