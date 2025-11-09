package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.143
	const number float64 = 2.00

	fmt.Println("PI value: ", PI)
	fmt.Println("Dividing number: ", math.Floor(PI/number))
}
