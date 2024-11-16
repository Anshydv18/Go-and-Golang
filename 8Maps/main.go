package main

import "fmt"

func areMapsEqual(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, val := range m1 {
		if m2[key] != val {
			return false
		}
	}
	return true
}

func main() {
	myarr := make(map[string]int)
	fmt.Println(myarr)
	m := map[string]int{
		"Alice": 1,
		"Bob":   2,
	}
	fmt.Println(m)
	myarr["Ansh"] = 10
	fmt.Println(myarr["Ansh"])

	age, exits := m["Alice"]
	if exits {
		fmt.Println(age)
	}
	//deleting elements
	delete(m, "Alice")
	// len of map is find as len(m)
	// and map is always a reference , changes in function in the map is reflected in the other

	//A map declared but not initialized is nil. Accessing elements in a nil map is safe, but adding elements will cause a runtime panic.
}
