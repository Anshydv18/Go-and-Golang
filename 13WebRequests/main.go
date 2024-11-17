package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//get call;
	urlpath := "https//:www.google.com/path/"
	response, err := http.Get(urlpath)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, er := ioutil.ReadAll(response.Body)
	if er != nil {
		panic(er)
	}
	fmt.Printf(string(body))

	///post path
	jsonData := []byte(
		`{
		"title":"hero",
		"body":"bar"
		}`,
	)

	res, err := http.Post(urlpath, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//we can made put, delete
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, urlpath, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-Type", "application/json")
	req.Header.Set("authorization", "Bearer Token")
	//we can adder header in the request
	Res, e := client.Do(req)
	if e != nil {
		panic(e)
	}
	defer Res.Body.Close()

}
