package main

import "fmt"
import "net/http"
import "encoding/json"

var url = "https://www.olx.co.id/api/relevance/v4/search?category=198&facet_limit=100&location=4000020&location_facet_limit=20&platform=web-desktop&query=cr%20v%202013&relaxedFilters=true&size=40&spellcheck=true&user=186274cf010x52b9387a"

type responsePayload struct {
	Data []item `json:data`
}

type item struct {
	Title string `json:title`
	Desc  string `json:description`
	Price price  `json:price`
}

type price struct {
	Value priceValue `json:value`
}

type priceValue struct {
	Display string `json:display`
}

func fetchItems() (*responsePayload, error) {
	var err error
	var client = &http.Client{}
	var data responsePayload

	request, err := http.NewRequest("GET", url, nil)

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

func main() {
	fmt.Println("hello")
	data, err := fetchItems()

	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	for _, each := range data.Data {
		fmt.Println(each.Title, each.Price.Value.Display)
		fmt.Println(each.Desc)
	}

	//fmt.Println(data)
}
