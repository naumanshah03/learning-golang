package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to Web Request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	responseDataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	responseHeader := response.Header
	fmt.Printf("Response header: %s\n", responseHeader)

	responseContent := string(responseDataBytes)
	fmt.Printf("Response content: %s\n", responseContent)
}
