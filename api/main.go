package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	post()
	url := "https://jsonplaceholder.typicode.com/posts"

	// basically takes request type, api endpoint and request body

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	// send api call
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()
	_, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr.Error())
	}

	// fmt.Println(string(body))

}

func post() {
	url := "https://jsonplaceholder.typicode.com/posts"
	testData := []byte(`{"title":"test title", "body":"this is the body"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(testData))

	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Add("Content-type", "application/json; charset=UTF-8")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)

	if readErr != nil {
		fmt.Println(readErr.Error())
	}

	fmt.Println(string(body))

}
