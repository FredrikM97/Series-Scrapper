package http

import (
	gb "Series-Scrapper/global"
	"Series-Scrapper/utils"
	"fmt"
	"regexp"
)

type myAnimeList struct {
	score    string
	rank     string
	episodes string
	info     string
	aired    string
	genre    string
	response string
	url      string
}

var MyAnimeList = myAnimeList{
	score:    `<span itemprop="ratingValue">((.|\n)*?)</span>`,
	rank:     `<span class="dark_text">Ranked:</span>\n  ((.|\n)*?)<sup>`,
	episodes: `<span class="dark_text">Episodes:</span>((.|\n)*?)</div>`,
	info:     `<meta property="og:description" content="((.|\n)*?)">`,
	aired:    `<span class="dark_text">Aired:</span>((.|\n)*?)</div>`,
	//Top =
	genre: `<span class="dark_text">Genres:</span>||</div>`,
	//Seasonal =
}
var siteVars = gb.WebsiteVars{
	URLTop:       `https://myanimelist.net/topanime.php?type=bypopularity`,
	URLSeason:    `https://myanimelist.net/anime/season`,
	RegexArticle: `<h2 id="anime">Anime</h2>(.|\n)*?</article>`,
	RegexAdress:  `<div class="picSurround di-tc thumb">([\s\S]*?)<a href="(https://myanimelist.net/anime/[0-9]*/([^"/]*))`,
	/*TODO: The problem is here*/
}

func (d myAnimeList) Search(searchURL string, contentName string) ([]string, bool) {
	/*
		Search for series on MAL
		Params: string
		Return:
		* bool: No found-> false
		* string:
	*/
	var err bool
	var params []string
	MyAnimeList.response, err = GetContent(searchURL)
	if err {
		return params, false
	}

	// Regex to find wanted info in data_response
	//articleRegex := `<h2 id="anime">Anime</h2>(.|\n)*?</article>`
	//addressRegex := `<div class="picSurround di-tc thumb">
	//<a href="https://myanimelist.net/anime/[0-9]*/([^"/]*)`

	re := regexp.MustCompile(siteVars.RegexArticle)
	queue := re.FindAllString(string(MyAnimeList.response), -1)[0] //First index is anime

	re = regexp.MustCompile(siteVars.RegexAdress)
	addresses := re.FindAllStringSubmatch(queue, -1)
	//fmt.Println("Current addresses", addresses)
	seriesMap := make(map[string]int)

	// Fix series name and check if they same as requested name
	var url string
	for index, info := range addresses {
		name := utils.Address2string(info[3])

		if name == contentName {
			url := info[2]

			MyAnimeList.response, _ = UpdateResponse(url)
			params := GetParameterValues(MyAnimeList)

			return params, true
		}
		seriesMap[name] = index
	}
	sortedMap := utils.OnKeyValue(seriesMap)

	// If series not found, recommend similar series

	if url == "" {
		fmt.Println("Could not find series! Did you mean: \n-------------------")
		for _, k := range sortedMap {
			fmt.Printf("%v: %v\n", k.Value, k.Key)
		}

	}

	return params, false
}
func (d myAnimeList) GetSeasonal() (string, bool) {
	//https://myanimelist.net/anime/season/2020/spring
	response, err := GetContent(siteVars.URLSeason)
	if err {
		return "", false
	}
	return response, true
}
func (d myAnimeList) GetTop() (string, bool) {

	response, err := GetContent(siteVars.URLTop)
	if err {
		return "", false
	}
	return response, true
}

func (d myAnimeList) GetScore() string {
	re := regexp.MustCompile(MyAnimeList.score)
	queue := re.FindAllStringSubmatch(string(MyAnimeList.response), -1)[0] //First index is anime
	return queue[1]
}
func (d myAnimeList) GetRank() string {
	re := regexp.MustCompile(MyAnimeList.rank)
	queue := re.FindAllStringSubmatch(string(MyAnimeList.response), -1)[0] //First index is anime
	return queue[1]
}
func (d myAnimeList) GetEpisodes() string {
	re := regexp.MustCompile(MyAnimeList.episodes)
	queue := re.FindAllStringSubmatch(string(MyAnimeList.response), -1)[0] //First index is anime
	return queue[1]
}
func (d myAnimeList) GetInfo() string {
	re := regexp.MustCompile(MyAnimeList.info)
	queue := re.FindAllStringSubmatch(string(MyAnimeList.response), -1)[0] //First index is anime
	return queue[1]
}
func (d myAnimeList) GetAired() string {
	re := regexp.MustCompile(MyAnimeList.aired)
	queue := re.FindAllStringSubmatch(string(MyAnimeList.response), -1)[0] //First index is anime
	return queue[1]
}
