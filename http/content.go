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
		data := `Score: ` + model.GetScore()
		params = append(params, data)
	}
	if gb.CommandMap.Rank {
		data := `Rank: ` + model.GetRank()
		params = append(params, data)
	}
	if gb.CommandMap.Episodes {
		data := `Episodes: ` + model.GetEpisodes()
		params = append(params, data)
	}
	if gb.CommandMap.Info {
		data := `Info: ` + model.GetInfo()
		params = append(params, data)
	}
	if gb.CommandMap.Aired {
		data := `Aired: ` + model.GetAired()
		params = append(params, data)
	}
	if gb.CommandMap.Top {
		top, _ := model.GetTop()
		data := `Top series: ` + top
		params = append(params, data)
	}

	return params
}
