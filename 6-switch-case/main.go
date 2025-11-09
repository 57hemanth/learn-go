package main

import "fmt"

func main() {
	digit := 6

	fmt.Println("Digit is:")

	switch digit {
	case 1:
		fmt.Print("One.")
	case 2:
		fmt.Print("Two.")
	case 3:
		fmt.Print("Three.")
	case 4:
		fmt.Print("Four.")
	case 5:
		fmt.Print("Five.")
	case 6:
		fmt.Print("Six.")
	case 7:
		fmt.Print("Seven.")
	case 8:
		fmt.Print("Eight.")
	case 9:
		fmt.Print("Nine.")
	default:
		fmt.Println("Invalid digit.")
	}
}
