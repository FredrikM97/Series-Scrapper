package myanimelist

import (
	"fmt"
	"regexp"
)

func search_MAL(search_url string) bool {
	/*
		Search for series on MAL
		Params: string
		Return:
		* bool: No found-> false
		* string:
	*/
	data_response, err := getContent(search_url)
	if err {
		return false
	}

	// Regex to find wanted info in data_response
	article_regex := `<h2 id="anime">Anime</h2>(.|\n)*?</article>`
	address_regex := `<div class="picSurround di-tc thumb">
    <a href="https://myanimelist.net/anime/[0-9]*/([^"/]*)`

	re := regexp.MustCompile(article_regex)
	result_queue := re.FindAllString(string(data_response), -1)[0] //First index is anime

	re = regexp.MustCompile(address_regex)
	result_addresses := re.FindAllStringSubmatch(result_queue, -1)

	seriesMap := make(map[string]int)

	// Fix series name and check if they same as requested name
	for index, serie_info := range result_addresses {
		name := address2string(serie_info)

		if name == command_map.name {
			command_map.url = serie_info[0]
			return true
		}
		seriesMap[name] = index
	}
	sorted_map := sort_on_keyValue(seriesMap)

	// If series not found, recommend similar series
	if command_map.url == "" {
		fmt.Println("Could not find series! Did you mean: \n-------------------")
		for _, k := range sorted_map {
			fmt.Printf("%v: %v\n", k.Value, k.Key)
		}

	}

	return true
}
