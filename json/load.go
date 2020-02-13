package json

import (
	jsonImp "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Open(db string, data interface{}) interface{} {
	/*
		Import data from database
	*/

	file, err := os.Open(db)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Successfully opened " + db)
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	jsonImp.Unmarshal(byteValue, &data)

	return data
}
