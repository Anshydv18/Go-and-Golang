package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buffer := bufio.NewReader(os.Stdin)
	input, err := buffer.ReadString('\n')
	if err != nil {
		fmt.Println("error in reading the input")
	}
	fmt.Println("the input is ", input)
}
