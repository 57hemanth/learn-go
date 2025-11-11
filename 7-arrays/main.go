package main

import (
	"fmt"
)

func main() {
	var a [5]int

	fmt.Println("A value: ", a)
	// Output: A value:  [0 0 0 0 0]

	a[4] = 20
	fmt.Println("A: ", a)

	// Length of the array
	fmt.Println("Length of A: ", len(a))

	b := [5]int{ 1, 2, 3, 4, 5}
	fmt.Println("B: ", b)

	c := [...]int{1, 4, 20, 40}
	fmt.Println("C: ", c)

	// Here we can define index to assign value and inbetween that index it will asign zeros.
	e := [...]int{3, 4: 2, 7}
	fmt.Println("E: ", e)
	// Output: E:  [3 0 0 0 2 7]

	// Two D Array
	var twoD [2][2]int
	for i := range 2 {
		for j := range 2 {
			twoD[i][j] = i * j
		}
	}
	fmt.Println("Two D: ", twoD)

	// Declaring 2D Array
	var twoDTwo = [2][3]int {
		{ 1, 2, 3},
		{ 4, 5, 6},
	}
	fmt.Println("Two D 2: ", twoDTwo)
}

