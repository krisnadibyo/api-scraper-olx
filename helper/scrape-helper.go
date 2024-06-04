package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var Url = "https://www.olx.co.id/api/relevance/v4/search?category=198&facet_limit=100&location_facet_limit=20&platform=web-desktop&relaxedFilters=true&size=40&spellcheck=true&user=186274cf010x52b9387a"

type ResponsePayload struct {
	Data []Item `json:"data"`
}

type Item struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"description"`
	Price Price  `json:"price"`
}

type Price struct {
	Value PriceValue `json:"value"`
}

type PriceValue struct {
	Display string `json:"display"`
}

func GenerateUrl(url string, keyword string, page int) string {
	location := "4000020" // Bekasi

	if page > 0 {
		return url + "&query=" + keyword + "&location=" + location + "&page=" + strconv.Itoa(page)
	}
	return url + "&query=" + keyword + "&location=" + location
}

func FetchItems(keyword string, page int) (*ResponsePayload, error) {
	var err error
	var client = &http.Client{}
	var data ResponsePayload

	request, err := http.NewRequest("GET", GenerateUrl(Url, keyword, page), nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func FetchData(keyword string) []Item {
	keyword = strings.Replace(keyword, " ", "%20", -1)
	data, err := FetchItems(keyword, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	data2, err := FetchItems(keyword, 1)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	data.Data = append(data.Data, data2.Data...)
	return data.Data
}

func FetchDatas(keywords []string) []Item {
	result := []Item{}
	for _, keyword := range keywords {
		keyword = strings.Replace(keyword, " ", "%20", -1)
		data, err := FetchItems(keyword, 0)
		if err != nil {
			fmt.Println("Error", err.Error())
		}

		data2, err := FetchItems(keyword, 1)
		if err != nil {
			fmt.Println("Error", err.Error())
		}
		result = append(result, data.Data...)
		result = append(result, data2.Data...)
	}
	return result
}
