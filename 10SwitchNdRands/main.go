package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the workweek.")
		fallthrough //fallthrough is used for continue in the case of switch case
	case "Friday":
		fmt.Println("Almost the weekend!")
	default:
		fmt.Println("It's a regular day.")
	}

	num := 10
	switch {
	case num > 0:
		fmt.Println("Positive number")
	case num < 0:
		fmt.Println("Negative number")
	default:
		fmt.Println("Zero")
	}

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	//in the random number ,the rand number can be deterministic because they use the same seed
	//in go we can cange the seed
	rand.Seed(time.Now().UnixNano())

	number := rand.Intn(4)
	fmt.Println(number)
}
