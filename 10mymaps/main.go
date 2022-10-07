package main

import "fmt"

func main() {
	fmt.Println("Learning maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages: ", languages)
	fmt.Println("JS short for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of languages: ", languages)

	// loops learning maps for golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
