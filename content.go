package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getContent(url string) (string, bool) {
	/*
		Get content on page
		@Params
		* url: url to page

		@return
		* string: info from site
		* bool: if true: Error
	*/

	data_response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: Could not fetch page!")
		return "", true
	}
	defer data_response.Body.Close()

	body_content, err := ioutil.ReadAll(data_response.Body)
	body_content2string := string(body_content)

	return body_content2string, false
}
