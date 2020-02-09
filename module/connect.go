package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func json_handler(db string, data interface{}) interface{} {
	/*
		Import data from database
	*/

	db_file, err := os.Open(db)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Successfully opened " + db)
	defer db_file.Close()

	byteValue, _ := ioutil.ReadAll(db_file)

	json.Unmarshal(byteValue, &data)

	return data
}
