package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in GoLang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, message := proAdder(2, 5, 7, 9, 2)
	fmt.Println("ProResult is: ", proResult)
	fmt.Println("ProMessage is: ", message)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namastey from GoLang")
}
