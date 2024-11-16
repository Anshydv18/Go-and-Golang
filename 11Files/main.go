package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Content for writing in the file"
	file, err := os.Create("./mytext.txt")
	if err != nil {
		fmt.Printf("error while creating the file")
		panic(err)
	}
	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileRead("./mytext.txt")
	fmt.Println("the len is ", len)
}

func fileRead(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("txt data inside the file is ", string(databyte))
}
