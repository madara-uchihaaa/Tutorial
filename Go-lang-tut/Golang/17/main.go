package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func getURL(port int) string {
	return fmt.Sprintf("http://localhost:%d/", port)
}

func postURL(port int) string {
	return fmt.Sprintf("http://localhost:%d/login", port)
}

func postFormURL(port int) string {
	return fmt.Sprintf("http://localhost:%d/postForm", port)
}

func getREQ() {
	url := getURL(3000)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	cntnt, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Status: ", response.Status)
	fmt.Println("Cntnt: ", string(cntnt))
}

func postREQ() {
	postUrl := postURL(3000)
	fmt.Println("URL: ", postUrl)

	reqBody := strings.NewReader(`{"username": "admin", "password": "admin"}`)

	res, err := http.Post(postUrl, "application/json", reqBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	cntnt, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Status: ", res.Status)
	fmt.Println("Cntnt: ", string(cntnt))
}

func postFormReq() {
	postFormUrl := postFormURL(3000)
	fmt.Println("URL: ", postFormUrl)

	data := url.Values{}
	data.Add("username", "admin")
	data.Add("password", "admin")
	data.Add("email", "xyz@mailcom")

	res, err := http.PostForm(postFormUrl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	cntnt, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Status: ", res.Status)
	fmt.Println("Cntnt: ", string(cntnt))
}

func main() {
	// getREQ()
	// postREQ()
	postFormReq()
}
