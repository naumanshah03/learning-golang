package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 30
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)  // ptr means memory address of the value stored
	fmt.Println("Value of actual pointer is ", *ptr) // *ptr means actual value

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)
}
