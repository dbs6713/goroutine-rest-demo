package main

// This program makes HTTP requests to three different URLs in serial.

import (
	"fmt"
	"net/http"
)

// Get a URL and output the status code response.
func get(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url, res.StatusCode)

}

func main() {
	get("http://bbc.co.uk")
	get("http://google.com")
	get("http://github.com")
}
