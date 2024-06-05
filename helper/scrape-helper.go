package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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

func FetchItems(url string, keyword string, page int) (*ResponsePayload, error) {
	var err error
	var client = &http.Client{}
	var data ResponsePayload

	request, err := http.NewRequest("GET", GenerateUrl(url, keyword, page), nil)

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

func FetchData(url string, keyword string) []Item {
	keyword = strings.Replace(keyword, " ", "%20", -1)
	data, err := FetchItems(url, keyword, 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	data2, err := FetchItems(url, keyword, 1)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	data.Data = append(data.Data, data2.Data...)
	return data.Data
}

func FetchDatas(url string, keywords []string) []Item {
	result := []Item{}
	for _, keyword := range keywords {
		keyword = strings.Replace(keyword, " ", "%20", -1)
		data, err := FetchItems(url, keyword, 0)
		if err != nil {
			fmt.Println("Error", err.Error())
		}

		data2, err := FetchItems(url, keyword, 1)
		if err != nil {
			fmt.Println("Error", err.Error())
		}
		result = append(result, data.Data...)
		result = append(result, data2.Data...)
	}
	return result
}
