package system

import (
	"fmt"
	"time"
)

func timeArithmetic() {
	bySec := time.Second / 5
	byMill := time.Millisecond * 200

	isSame := bySec == byMill
	fmt.Printf("Is same %t\n", isSame)
	fmt.Printf("1: %d 2: %d\n", bySec.Milliseconds(), byMill.Milliseconds())
}
