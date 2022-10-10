package main

import "fmt"

func main() {
	fmt.Println("Weclome to loops in GoLang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for each loop
	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at learncodeonline.in")
}
