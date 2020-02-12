package main

import (
	"fmt"
	"regexp"
)

func SearchMAL(searchURL string) bool {
	/*
		Search for series on MAL
		Params: string
		Return:
		* bool: No found-> false
		* string:
	*/
	resp, err := GetContent(searchURL)
	if err {
		return false
	}

	// Regex to find wanted info in data_response
	articleRegex := `<h2 id="anime">Anime</h2>(.|\n)*?</article>`
	addressRegex := `<div class="picSurround di-tc thumb">
    <a href="https://myanimelist.net/anime/[0-9]*/([^"/]*)`

	re := regexp.MustCompile(articleRegex)
	queue := re.FindAllString(string(resp), -1)[0] //First index is anime

	re = regexp.MustCompile(addressRegex)
	addresses := re.FindAllStringSubmatch(queue, -1)

	seriesMap := make(map[string]int)

	// Fix series name and check if they same as requested name
	for index, info := range addresses {
		name := Address2string(info)

		if name == commandMap.name {
			commandMap.url = info[0]
			return true
		}
		seriesMap[name] = index
	}
	sortedMap := SortOnKeyValue(seriesMap)

	// If series not found, recommend similar series
	if commandMap.url == "" {
		fmt.Println("Could not find series! Did you mean: \n-------------------")
		for _, k := range sortedMap {
			fmt.Printf("%v: %v\n", k.Value, k.Key)
		}

	}

	return true
}
