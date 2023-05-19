package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	const myURL = "https://dsa101.onrender.com/get-chapters"

	res, err := http.Get(myURL)
	if err != nil {
		log.Fatalf("Error occurred while making the HTTP request: %v", err)
	}
	defer res.Body.Close()

	fmt.Println("Status code:", res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error occurred while reading the response body: %v", err)
	}

	fmt.Println(string(body))
}