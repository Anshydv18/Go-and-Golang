package main

import "fmt"

func main() {
	var num int = 2
	var ptr *int = &num
	//& is  used to  reference the value and * is used to deferencing
	fmt.Println(*ptr, &ptr, &num)

	*ptr = *ptr * 2
	res := *ptr * 2
	fmt.Println(*ptr, res)
}
