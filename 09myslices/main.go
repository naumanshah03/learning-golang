package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video of slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 945
	highScores[2] = 789
	highScores[3] = 121
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 456)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remve a value from slices bsaed on index

	var courses = []string{"reactjs", "javascript", "swift", "golang", "python", "java"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
