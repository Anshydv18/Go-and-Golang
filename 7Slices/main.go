package main

import "fmt"

//comparing equality of slices

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	var s []int                    //A slice with no elements
	fmt.Println(s, len(s), cap(s)) //output :[] 0 0

	d := make([]int, 4)
	fmt.Println(d)

	arr := [5]int{1, 2, 3, 4, 5}
	newvar := arr[1:4]

	//acccessing and modifying is same like array
	fmt.Println(newvar)

	myarr := append(newvar, 1, 2, 3)
	fmt.Println(myarr)
}
