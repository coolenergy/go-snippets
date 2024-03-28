package conversions

import (
	"fmt"
	"strconv"
)

func parseNumbers() {

	const intString = "40"
	const hexString = "ff"    // 255
	const binString = "00011" // 3
	const floatString = "26.52"

	intInt, err := strconv.Atoi(intString)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Atoi(): %d\n", intInt)

	hexInt, err := strconv.ParseInt(hexString, 16, 32)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ParseInt base 16: %d\n", hexInt)

	binNumber, err := strconv.ParseInt(binString, 2, 32)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ParseInt base 2: %d\n", binNumber)

	floatNumber, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ParseInt: %.4f\n", floatNumber)
}
