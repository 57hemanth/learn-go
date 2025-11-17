package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("Uninitialized s: ", s, s == nil, len(s) == 0)

	s = make([]string, 3, 5)
	fmt.Println("S: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("S: ", s)

	s = append(s, "d")
	fmt.Println("S: ", s)

	s = append(s, "e")
	fmt.Println("S: ", s)

	s = append(s, "f")
	fmt.Println("S: ", s)

	s = append(s, "g")
	fmt.Println("S: ", s, "Len: ", len(s), "Cap: ", cap(s))

	// Creating a copy
	var c = make([]string, len(s))
	copy(c, s)
	fmt.Println("C: ", c, "Len: ", len(c), "Cap: ", cap(c)) // Here, capacity will be equal to the length

	// Slicing
	fmt.Println("1 to 3", s[1:3]) // Excluding '3'
	
	var t1 = []string{"a", "b", "c"}
	fmt.Println("T1: ", t1, "Len: ", len(t1), "Cap: ", cap(t1))

	var t2 = []string{"a", "b", "c"}
	fmt.Println("T2: ", t2, "Len: ", len(t2), "Cap: ", cap(t2))

	// Slices Equal function
	if slices.Equal(t1, t2) {
		fmt.Println("T1 and T2 are equal.")
	}

	// Two Dimensional Slice
	twoD := make([][]int, 3)

	for i := range 3 {
		innerLoop := i + 1
		twoD[i] = make([]int, innerLoop)
		for j := range innerLoop {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("Two D: ", twoD)
}
