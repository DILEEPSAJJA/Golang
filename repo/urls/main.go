package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Handling URL in golang")
	fmt.Println(myURL)

	result, _ := url.Parse(myURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)

	qparams := result.Query()
	fmt.Println("The Type of the Query params is : ", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Param is : " , val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}