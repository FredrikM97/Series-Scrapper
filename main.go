package main
import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/mitchellh/mapstructure"
	"flag"
	//"net/http"
	//"strings"
)

// Global variables
var DATABASE = "database/"

// Structs
type Sites struct {
    Sites 	[]Site 		`json:"sites"`
}

type Site struct {
	Name	string		`json:"name"`
	Url   	string 		`json:"url"`
	Genre  	[]string 	`json:"genre"`
}

var (
	name, seasonal, url					string 
	score, rank, episodes,info, aired	bool

)



func main() {
	/*
	Init system from here
	*/
	fmt.Println("Starting system!")
	var sites Sites

	data := json_handler(DATABASE + "sites.json", sites)
	mapstructure.Decode(data, &sites)

	fmt.Println("\nData: " +  fmt.Sprint(sites))
	setFlags()
}

func setFlags() {	
	/*
	Handles params for shell, TODO: Fix so commands stay in database instead (easier to organize)
	*/
	flag.StringVar(&name, "name", "", "Write a name of series")
	flag.StringVar(&seasonal, "seasonal", "", "Format: <SEASON> <YEAR>, blank gives the current season")
	flag.BoolVar(&score, "score", false, "Get score of series")
	flag.BoolVar(&rank, "rank", false, "Get rank of series")
	flag.BoolVar(&episodes, "episodes", false, "Get number of episodes")
	flag.BoolVar(&info, "info", false, "Get information of series")
	flag.BoolVar(&aired, "aired", false, "Get aired date of series")
        
	flag.Parse()
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

	fmt.Println("\nData: " +  fmt.Sprint(data))
	return data
}
func print_db(sites Sites){
	/*
	Print info from the database
	*/
	for i := 0; i < len(sites.Sites); i++ {
		fmt.Println("Name: " + sites.Sites[i].Name)
		fmt.Println("Url: " + sites.Sites[i].Url)
		fmt.Println("Genres: " + fmt.Sprint(sites.Sites[i].Genre))
	}
}
func getContent(url string) (string, bool) {
	/*
	Get content on page
	@Params
	* url: url to page

	@return
	* string: info from site
	* bool: if match found or not
	*/
	/*
	data_resp, err := http.GET(url)

	if err != nil {
		fmt.Println("Error: Could not fetch page!")
		return "", true
	}
	defer data_resp.Close()
	*/
	return "",false

}

func getSeasonal(){

}
func getTopTen(){

}

func Search() {
	/*
	Search on site based on parameters from shell
	*/
}