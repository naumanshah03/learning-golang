package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")

	const myUrl = "http://localhost:8000/get"
	// PerformGetRequest(myUrl)
	// PostJsonRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")

}

func PerformGetRequest(myUrl string) {

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println((string(content)))
}

func PostJsonRequest(myUrl string) {
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Let's go with GoLang",
			"price" : 0,
			"platform": "learncodeonline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest(myUrl string) {
	// fake form payload

	data := url.Values{}
	data.Add("firstName", "Nauman")
	data.Add("lastName", "Shah")
	data.Add("email", "nauman@go.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
