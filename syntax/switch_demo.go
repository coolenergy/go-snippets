package syntax

import "fmt"

func SwitchDemo() {
	day := "monday"

	switch day {
	case "monday":
		example := 2
		fmt.Println("Monday!!!!!!")
		fmt.Println(example)
		// notice we do not have to call break
	case "tuesday":
		fmt.Println("Tuesday!!!!")
		// Notice how the app still compiles, in Java for example it would not, stating the var is already defined
		notShadowedVar := 3
		fmt.Println(notShadowedVar)
	default:
		fmt.Println("Unexpected value")
	}

	number := 2

	switch number {
	case 1:
		fmt.Println("Monday!!!!!!")

	case 2:
		fmt.Println("Number 2 is also cool")
	}
}
