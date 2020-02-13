package main

import (
	gb "Series-Scrapper/global"
	"Series-Scrapper/http"
	"Series-Scrapper/json"
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

/*
Structs
*/

type sites struct {
	Sites []site `json:"sites"`
}

type site struct {
	Name   string            `json:"name"`
	URL    string            `json:"url"`
	Search string            `json:"search"`
	Genre  map[string]string `json:"genre"`
}

// Convert map value to function, return bool

/*
Global var
*/
const DATABASE = "database/"

var SitesAvailable = map[string]gb.Website{
	"myanimelist":  http.MyAnimeList,
	"myanimelist2": http.MyAnimeList,
}

var sitesInfo sites

func main() {
	/*
		Init system from here
	*/
	fmt.Println("Starting system!")
	setFlags()

	data := json.Open(DATABASE+"sites.json", sitesInfo)
	mapstructure.Decode(data, &sitesInfo)

	printDB(sitesInfo)
	// Check input
	if gb.CommandMap.Name != "" {
		check := checkParams()
		if check {
			site := sitesInfo.Sites[0] // TODO: Change so we check all availbile sites
			success := Search(site)
			if success {
				//getparameterValues()
			}

		}
	} else if gb.CommandMap.Seasonal != "" {

	}
}
func checkParams() bool {
	/*
		Check if any parameters is given, if not, return boolcheck false
		@Params: Commands

		@Return: bool
	*/
	paramExists := false
	r := reflect.ValueOf(gb.CommandMap)
	for i := 0; i < r.NumField(); i++ {
		//Get value of param
		f := r.Field(i)

		if f.Kind() == reflect.Bool {
			if reflect.ValueOf(true).Bool() == f.Bool() {
				paramExists = true
				return paramExists
			}
		}

	}
	return paramExists
}

func setFlags() {
	/*
		Handles params for shell, TODO: Fix so commands stay in database instead (easier to organize)
	*/

	flag.StringVar(&gb.CommandMap.Name, "name", "", "Write a name of series")
	flag.StringVar(&gb.CommandMap.Seasonal, "seasonal", "", "Format: <SEASON> <YEAR>, blank gives the current season")
	flag.BoolVar(&gb.CommandMap.Score, "score", false, "Get score of series")
	flag.BoolVar(&gb.CommandMap.Rank, "rank", false, "Get rank of series")
	flag.BoolVar(&gb.CommandMap.Episodes, "episodes", false, "Get number of episodes")
	flag.BoolVar(&gb.CommandMap.Info, "info", false, "Get information of series")
	flag.BoolVar(&gb.CommandMap.Aired, "aired", false, "Get aired date of series")

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

func Search(siteInfo site) bool {
	/*
		Search on site based on parameters from shell
		TODO: Should check if parameters is given or not too
	*/
	// TODO: Handle multiple different sites

	if _, ok := SitesAvailable[siteInfo.Name]; !ok {
		return false
	}

	var genre string
	if val, ok := siteInfo.Genre[gb.CommandMap.Genre]; ok {
		genre = siteInfo.Genre[val]
	} else {
		genre = siteInfo.Genre["0"]
	}
	r := strings.NewReplacer("*genre*", genre)
	search := r.Replace(siteInfo.Search)

	searchURL := siteInfo.URL + search + strings.ToLower(gb.CommandMap.Name)
	params, success := SitesAvailable[siteInfo.Name].Search(searchURL, gb.CommandMap.Name)

	print_params(params)
	return success

}
func print_params(params []string) {
	for item := range params {
		fmt.Println(item)
	}

}
