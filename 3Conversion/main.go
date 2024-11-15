package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error while Reading the input", err)
	}

	number, e := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if e != nil {
		fmt.Println("error while converting into float", e)
	}
	fmt.Println("after converting the number ", number)
}
