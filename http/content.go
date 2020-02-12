package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetContent(url string) (string, bool) {
	/*
		Get content on page
		@Params
		* url: url to page

		@return
		* string: info from site
		* bool: if true: Error
	*/

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: Could not fetch page!")
		return "", true
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return toString(body), false
}

func toString(data []byte) string { return string(data) }