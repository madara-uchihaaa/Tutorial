package main

import (
	"fmt"
	"net/url"
)

const URL = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	res, err := url.Parse(URL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", res)
	fmt.Println("Response is: ", res)

	fmt.Println("Scheme: ", res.Scheme)
	fmt.Println("Host: ", res.Host)
	fmt.Println("Path: ", res.Path)
	fmt.Println("Port: ", res.Port())
	fmt.Println("RawQuery: ", res.RawQuery)
	fmt.Println("Query(): ", res.Query())
	fmt.Println("Query(): ", res.Query().Get("coursename"))

	qparams := res.Query()
	fmt.Printf("Type of qparams: %T\n", qparams)
	// Output: Type of qparams: url.Values means url.Values is a map[string][]string

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=aman",
	}
	fmt.Println(partsOfUrl.String())
}