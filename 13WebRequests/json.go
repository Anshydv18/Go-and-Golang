package main

import (
	"encoding/json"
	"fmt"
)

// Define the struct to match the JSON structure
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

func jsonData() {
	// JSON data
	jsonData := `{"name": "John Doe", "age": 30, "email": "john.doe@example.com", "active": true}`

	// Declare a variable of type Person
	var person Person

	// Unmarshal JSON into the struct
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Access struct fields
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Email:", person.Email)
	fmt.Println("Active:", person.Active)
}
