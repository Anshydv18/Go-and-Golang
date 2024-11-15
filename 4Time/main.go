package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	//if we want to format the time into a proper manner we can use it
	createdeDate := time.Date(2026, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdeDate)
}
