package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Makes HTTP GET request on the given url
// Returns HTML and error (if any)
func getHtmlFromUrl(url string) (body string, statusCode int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("1 ERROR!!!1")
		return
	}
	// Status Code 200
	statusCode = resp.StatusCode

	defer resp.Body.Close()

	bytebody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("2 ERROR!!!1")
	}
	// Converting []byte to string
	body = string(bytebody)
	return
}

func main() {
	fmt.Println("Fetching HTML")
	html, statusCode, err := getHtmlFromUrl("https://google.com")
	if err != nil {
		fmt.Println("3 ERROR!!!1")
	} else {
		fmt.Println(statusCode, html)
	}
}
