package main

import (
	"Series-Scrapper/json"
	"Series-Scrapper/url"
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

/*
Structs
*/
type commands struct {
	name, seasonal, url, genre              string
	score, rank, episodes, info, aired, top bool // One of these bools must be true
}
type results struct {
	url                                string
	seasonal, genre, top               []string
	score, rank, episodes, info, aired string
}

type sites struct {
	Sites []site `json:"sites"`
}

type site struct {
	Name   string            `json:"name"`
	URL    string            `json:"url"`
	Search string            `json:"search"`
	Genre  map[string]string `json:"genre"`
}
type Map2func func(string) bool

// Convert map value to function, return bool

/*
Global var
*/
var DATABASE = "database/"
var commandMap = new(commands)
var resultMap = new(results)
var sitesInfo sites
var sitesAvailable = map[string]Map2func{
	"myanimelist": url.SearchMAL,
}

func main() {
	/*
		Init system from here
	*/
	fmt.Println("Starting system!")
	setFlags()

	data := json.LoadFile(DATABASE+"sites.json", sitesInfo)
	mapstructure.Decode(data, &sitesInfo)

	printDB(sitesInfo)
	// Check input
	if commandMap.name != "" {
		check := checkParams()
		fmt.Println("Is dis true=", commandMap.name)
		if check {
			site := sitesInfo.Sites[0] // TODO: Change so we check all availbile sites
			success := Search(site)
			if success {
				//getparameterValues(enabledParams)
			}

		}
	} else if commandMap.seasonal != "" {

	}
}
func checkParams() bool {
	/*
		Check if any parameters is given, if not, return boolcheck false
		@Params: Commands

		@Return: bool
	*/
	paramExists := false
	r := reflect.ValueOf(commandMap).Elem()
	for i := 0; i < r.NumField(); i++ {

		//Get value of param
		f := r.Field(i)
		if f.Kind() == reflect.Bool && reflect.ValueOf(true).Bool() == f.Bool() {
			paramExists = true
			continue
		}

	}
	return paramExists
}
func getparameterValues(enabledParams []bool) {
	/*
		Get data for parameters
	*/
	//score, rank, episodes, info, aired
	for _, item := range enabledParams {
		fmt.Println(item)

	}
}
func setFlags() {
	/*
		Handles params for shell, TODO: Fix so commands stay in database instead (easier to organize)
	*/

	flag.StringVar(&commandMap.name, "name", "", "Write a name of series")
	flag.StringVar(&commandMap.seasonal, "seasonal", "", "Format: <SEASON> <YEAR>, blank gives the current season")
	flag.BoolVar(&commandMap.score, "score", false, "Get score of series")
	flag.BoolVar(&commandMap.rank, "rank", false, "Get rank of series")
	flag.BoolVar(&commandMap.episodes, "episodes", false, "Get number of episodes")
	flag.BoolVar(&commandMap.info, "info", false, "Get information of series")
	flag.BoolVar(&commandMap.aired, "aired", false, "Get aired date of series")

	flag.Parse()
}

func printDB(sitesInfo sites) {
	/*
		Print info from the database
	*/
	fmt.Println("Sites:")
	for i := 0; i < len(sitesInfo.Sites); i++ {
		fmt.Println("Name: " + sitesInfo.Sites[i].Name)
		fmt.Println("Url: " + sitesInfo.Sites[i].URL)
		fmt.Println("Genres: " + fmt.Sprint(sitesInfo.Sites[i].Genre))
	}
}

func getSeasonal() {

}
func getTopTen() {

}

func getDetails() {

}

func Search(siteInfo site) bool {
	/*
		Search on site based on parameters from shell
		TODO: Should check if parameters is given or not too
	*/
	// TODO: Handle multiple different sites

	if _, ok := sitesAvailable[siteInfo.Name]; !ok {
		return false
	}

	var genre string
	if val, ok := siteInfo.Genre[commandMap.genre]; ok {
		genre = siteInfo.Genre[val]
	} else {
		genre = siteInfo.Genre["0"]
	}
	r := strings.NewReplacer("*genre*", genre)
	search := r.Replace(siteInfo.Search)

	searchURL := siteInfo.URL + search + strings.ToLower(commandMap.name)
	return sitesAvailable[siteInfo.Name](searchURL)

}
