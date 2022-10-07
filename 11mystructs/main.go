package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inhertance in golang; No super or parent

	nauman := User{"Nauman", "nauman@go.dev", true, 18}
	fmt.Println(nauman)
	fmt.Printf("Nauman details are: %+v\n", nauman)
	fmt.Printf("Name is %v and email is %v", nauman.Name, nauman.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
