package http

import (
	gb "Series-Scrapper/global"

	"fmt"
	"io/ioutil"
	httpImp "net/http"
)

func GetContent(url string) (string, bool) {
	/*
		Get content on page
		@Params
		* url: url to page

		@return
		* string: info from site
		* bool: if true: Error
	*/

	resp, err := httpImp.Get(url)

	if err != nil {
		fmt.Println("Error: Could not fetch page!")
		return "", true
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return toString(body), false
}

func UpdateResponse(url string) (string, bool) {
	response, err := GetContent(url)
	if err {
		return "", false
	}
	return response, true
}
func toString(data []byte) string { return string(data) }

func GetParameterValues(model gb.Website) []string {
	/*
		Get data for parameters
	*/
	//score, rank, episodes, info, aired
	var params []string
	if gb.CommandMap.Score {
		params = append(params, model.GetScore())
	}
	if gb.CommandMap.Rank {
		params = append(params, model.GetRank())
	}
	if gb.CommandMap.Episodes {
		params = append(params, model.GetEpisodes())
	}
	if gb.CommandMap.Info {
		params = append(params, model.GetInfo())
	}
	if gb.CommandMap.Aired {
		params = append(params, model.GetAired())
	}
	if gb.CommandMap.Top {
		top, _ := model.GetTop()
		params = append(params, top)
	}

	return params
}
