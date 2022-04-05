package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghsbf23"

func main() {
	fmt.Println("Handling URLS in Golang")
	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Ankit",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
