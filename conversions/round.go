package conversions

import (
	"fmt"
	"math"
)

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func roundSnippet() {
	floatNumber := 9021.2
	rounded := Round(floatNumber, 1.0)
	fmt.Printf("From %f to %f\n", floatNumber, rounded)

	floatNumber = 3.5
	rounded = Round(floatNumber, 1.0)
	fmt.Printf("From %f to %f\n", floatNumber, rounded)

	floatNumber = 780.2
	rounded = Round(floatNumber, 1000.0)
	fmt.Printf("From %f to %f\n", floatNumber, rounded)

	floatNumber = 40.0
	rounded = Round(floatNumber, 100.0)
	fmt.Printf("From %f to %f\n", floatNumber, rounded)

}
