package main

import (
	"fmt"
)

func main() {
	i := 1

	for {
		fmt.Println("Loops started...")
		break
	}

	for i <= 3 {
		fmt.Println("I value: ", i)
		i += 1
	}

	for j := 0; j < 10; j++ {
		fmt.Println("J value: ", j)
	}

	for k := range 10 {
		fmt.Println("K value: ", k)
	}
}
