package main

import (
	"fmt"
	"net/url"
)

func main() {
	Url := "https://example.com/path?query=hello&lang=en#section"
	parsedUrl, err := url.Parse(Url)
	if err != nil {
		fmt.Printf("error while parsing the url")
	}

	fmt.Println("scheme", parsedUrl.Scheme)
	fmt.Println("host", parsedUrl.Host)
	fmt.Println("Path ", parsedUrl.Path)
	fmt.Println("port is", parsedUrl.Port())
	fmt.Println("query is", parsedUrl.RawQuery)

	queryParams := parsedUrl.Query()
	fmt.Println("Query :", queryParams.Get("query"))
	fmt.Println("Lang :", queryParams.Get("lang"))

	baseUrl := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "path",
	}

	query := url.Values{}
	query.Add("query", "hello")
	query.Add("lang", "en")
	baseUrl.RawQuery = query.Encode()
	fmt.Println("Constructed url", baseUrl.String())
}
