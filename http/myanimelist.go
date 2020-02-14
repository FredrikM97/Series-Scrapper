package http

import (
	"Series-Scrapper/utils"
	"fmt"
	"regexp"
	"strings"
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

type vars struct {
	URLTop        string
	URLSeason     string
	searchArticle string
	searchAdress  string
	topArticle    string
	topAddress    string
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
var siteVars = vars{
	URLTop:        `https://myanimelist.net/topanime.php?type=bypopularity`,
	URLSeason:     `https://myanimelist.net/anime/season`,
	searchArticle: `<h2 id="anime">Anime</h2>(.|\n)*?</article>`,
	searchAdress:  `<div class="picSurround di-tc thumb">([\s\S]*?)<a href="(https://myanimelist.net/anime/[0-9]*/([^"/]*))`,

	topArticle: `<div class="anime-header">TV \(New\)</div>(.|\n)*?<div class="anime-header">`,
	topAddress: `<p class="title-text">([\s\S]*?)<a href="(https://myanimelist.net/anime/[0-9]*/([^"/]*))"`,
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
	var url string
	var seriesMap = make(map[string]int)

	re := regexp.MustCompile(siteVars.searchArticle)
	queue := re.FindAllString(string(MyAnimeList.response), -1)[0]

	re = regexp.MustCompile(siteVars.searchAdress)
	addresses := re.FindAllStringSubmatch(queue, -1)

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

	// If series not found, recommend similar series

	if url == "" {
		fmt.Println("Could not find series! Did you mean: \n-------------------")
		for _, k := range utils.OnKeyValue(seriesMap) {
			fmt.Printf("%v: %v\n", k.Value, k.Key)
		}

	}

	return params, false
}
func (d myAnimeList) GetSeasonal(data string) ([]string, bool) {
	//https://myanimelist.net/anime/season/2020/spring
	var err bool
	var params []string
	// If no season given then default to current season
	if data == "" {
		MyAnimeList.response, err = GetContent(siteVars.URLSeason)
	} else {
		splitData := strings.SplitN(data, " ", -1)
		season := splitData[0]
		year := splitData[1]

		MyAnimeList.response, err = GetContent(siteVars.URLSeason + "/" + year + "/" + season)
	}

	if err {
		return params, false
	}
	re := regexp.MustCompile(siteVars.topArticle)
	queue := re.FindAllString(string(MyAnimeList.response), -1)[0]

	re = regexp.MustCompile(siteVars.topAddress)
	addresses := re.FindAllStringSubmatch(queue, -1)
	//fmt.Println("Current addresses", addresses)
	seriesMap := make(map[string]int)

	for index, info := range addresses {
		name := utils.Address2string(info[3])
		seriesMap[name] = index
	}
	sortedMap := utils.OnKeyValue(seriesMap)

	var vals string
	for _, k := range sortedMap {
		vals = vals + "\n" + " " + k.Key
		//fmt.Printf("%v: %v\n", k.Value, k.Key)
	}
	params = append(params, vals)
	return params, true
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
