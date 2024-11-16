package main

import "fmt"

type person struct {
	Name string
	Age  int
	City string
	address
}

type address struct {
	Village  string `json:"name"`
	PostCode int    `json:"age"`
}

func main() {
	Ansh := person{
		Name: "Ansh",
		Age:  23,
		City: "Akbarpur",
		address: address{
			Village:  "Haidrabad",
			PostCode: 224151,
		},
	}
	fmt.Println(Ansh)
	// if we send the the struct  as reference, if we make changes in one file of function it will reflect the changes into the another function

}

func updatePerson(p *person) {
	p.Age = 40 // Changes are reflected outside the function
}
