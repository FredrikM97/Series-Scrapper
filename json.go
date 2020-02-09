package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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
