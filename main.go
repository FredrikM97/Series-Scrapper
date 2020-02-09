package main

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

/*
Structs
*/
type Commands struct {
	name, seasonal, url                string
	genre                              string
	score, rank, episodes, info, aired bool
}

type Sites struct {
	Sites []Site `json:"sites"`
}

type Site struct {
	Name   string            `json:"name"`
	Url    string            `json:"url"`
	Search string            `json:"search"`
	Genre  map[string]string `json:"genre"`
}

// Convert map value to function, return bool
type Map2func func(string) bool

/*
Global var
*/
var DATABASE = "database/"
var command_map = new(Commands)
var sites Sites
var available_sites = map[string]Map2func{
	"myanimelist": search_MAL,
}

func main() {
	/*
		Init system from here
	*/
	fmt.Println("Starting system!")
	setFlags()

	data := json_handler(DATABASE+"sites.json", sites)
	mapstructure.Decode(data, &sites)

	print_db(sites)

	// Check input
	if command_map.name != "" {
		bool_check := checkParams()
		if bool_check {
			site := sites.Sites[0] // TODO: Change so we check all availbile sites
			success := Search(site)

			if success {

			}

		}
	} else if command_map.seasonal != "" {

	}
}
func checkParams() bool {
	/*
		Check if any parameters is given, if not, return boolcheck false
		@Params: Commands

		@Return: bool
	*/
	r := reflect.ValueOf(command_map).Elem()
	for i := 0; i < r.NumField(); i++ {

		//Get value of param
		f := r.Field(i)

		if f.Kind() == reflect.Bool && reflect.ValueOf(true).Bool() == f.Bool() {
			return true
		}
	}
	return false
}
func setFlags() {
	/*
		Handles params for shell, TODO: Fix so commands stay in database instead (easier to organize)
	*/

	flag.StringVar(&command_map.name, "name", "", "Write a name of series")
	flag.StringVar(&command_map.seasonal, "seasonal", "", "Format: <SEASON> <YEAR>, blank gives the current season")
	flag.BoolVar(&command_map.score, "score", false, "Get score of series")
	flag.BoolVar(&command_map.rank, "rank", false, "Get rank of series")
	flag.BoolVar(&command_map.episodes, "episodes", false, "Get number of episodes")
	flag.BoolVar(&command_map.info, "info", false, "Get information of series")
	flag.BoolVar(&command_map.aired, "aired", false, "Get aired date of series")

	flag.Parse()
}

func print_db(sites Sites) {
	/*
		Print info from the database
	*/
	fmt.Println("Sites:")
	for i := 0; i < len(sites.Sites); i++ {
		fmt.Println("Name: " + sites.Sites[i].Name)
		fmt.Println("Url: " + sites.Sites[i].Url)
		fmt.Println("Genres: " + fmt.Sprint(sites.Sites[i].Genre))
	}
}

func getSeasonal() {

}
func getTopTen() {

}

func getDetails() {

}

func Search(site Site) bool {
	/*
		Search on site based on parameters from shell
		TODO: Should check if parameters is given or not too
	*/
	// TODO: Handle multiple different sites

	if _, ok := available_sites[site.Name]; !ok {
		return false
	}

	var genre string
	if val, ok := site.Genre[command_map.genre]; ok {
		genre = site.Genre[val]
	} else {
		genre = site.Genre["0"]
	}
	r := strings.NewReplacer("*genre*", genre)
	search := r.Replace(site.Search)

	search_url := site.Url + search + strings.ToLower(command_map.name)
	return available_sites[site.Name](search_url)

}

func Address2string(address []string) string {
	/*
		Extract name from list to string and remove parameters that shouldnt be included
		Params:	[]string
		Return: string
	*/
	r := strings.NewReplacer("_", " ", "  ", " ")
	serie_address := address[1]
	old_address := ""

	for old_address != serie_address {
		old_address = serie_address
		serie_address = r.Replace(serie_address)
	}

	serie_address = strings.ToLower(serie_address)
	return serie_address
}
