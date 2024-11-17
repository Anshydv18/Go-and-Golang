package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type person struct {
	Name string
	Age  int
	Tags []string
}

func Data() {
	data := []person{
		{"Ansh", 12, []string{"A", "B"}},
		{"Dev", 14, []string{"c", "d"}},
	}

	finaljson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(finaljson))
}

func Form() {
	path := "abc.com"
	data := url.Values{}
	data.Add("Name", "Ansh")
	data.Add("address", "gurugram")

	res, err := http.PostForm(path, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(content))
}
