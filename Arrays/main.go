package main

import (
	"fmt"
	"sort"
)

func main() {
	//initialising the array
	var arr [5]int //all elements are initialised to 0
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [5]int{1, 2}                // remaining elements are initialized to 0
	arr4 := [...]int{2, 4, 12, 3, 5, 6} //the size is determined at compilar time
	fmt.Println(arr, arr2, arr3, arr4, len(arr4))

	//accessing elements
	arr3[3] = 50
	fmt.Println(arr3)

	//iterating on the array
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	for idx, value := range arr4 {
		fmt.Printf("%d %d \n", idx, value)
	}
	//copying the array

	newArr := arr4
	sort.Ints(newArr[:])
	fmt.Println(newArr)

}
