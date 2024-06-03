package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var url = "https://www.olx.co.id/api/relevance/v4/search?category=198&facet_limit=100&location_facet_limit=20&platform=web-desktop&relaxedFilters=true&size=40&spellcheck=true&user=186274cf010x52b9387a"

type responsePayload struct {
	Data []item `json:data`
}

type item struct {
	ID    string `json:id`
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

type Token struct {
	PrivateKey   string `json:"private_key"`
	PrivateKeyID string `json:"private_key_id"`
	Email        string `json:"client_email"`
}

func GenerateUrl(url string, keyword string, page int) string {
	location := "4000020" // Bekasi

	if page > 0 {
		return url + "&query=" + keyword + "&location=" + location + "&page=" + strconv.Itoa(page)
	}
	return url + "&query=" + keyword + "&location=" + location
}

func fetchItems(keyword string, page int) (*responsePayload, error) {
	var err error
	var client = &http.Client{}
	var data responsePayload

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

func fetchData() []item {
	data, err := fetchItems("brio%202013", 0)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	data2, err := fetchItems("brio%202013", 1)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	data.Data = append(data.Data, data2.Data...)
	return data.Data
}

func readTokenFile() Token {
	jsonFile, err := os.Open("service-account.json")
	if err != nil {
		log.Fatalf("Unable to read token file: %v", err)
	}

	defer jsonFile.Close()

	token := Token{}
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&token); err != nil {
		log.Fatalf("Unable to parse token file: %v", err)
	}

	return token
}

func setupGsheet() (*sheets.Service, error) {
	token := readTokenFile()

	conf := &jwt.Config{
		Email:        token.Email,
		PrivateKey:   []byte(token.PrivateKey),
		PrivateKeyID: token.PrivateKeyID,
		TokenURL:     "https://oauth2.googleapis.com/token",
		Scopes: []string{
			"https://www.googleapis.com/auth/spreadsheets",
		},
	}

	client := conf.Client(context.Background())

	// Create a service object for Google sheets
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
		return nil, err
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	return srv, nil
}

func AppendRow(srv *sheets.Service, spreadsheetId string, title string, price string) {

	values := &sheets.ValueRange{
		Values: [][]interface{}{{
			title,
			price,
		}},
	}

	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, "Sheet1!A:F", values).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		log.Fatalf("Unable to insert data to sheet: %v", err)
	}
}

func main() {
	spreadsheetId := "13QBSQflOCao6HYJucZ02YnB8Uq5Q9l4cVEN0ab1VOXw"

	srv, _ := setupGsheet()
	cars := fetchData()
	for _, each := range cars {
		title := each.Title
		price := each.Price.Value.Display
		price = price[3:]
		price = strings.Replace(price, ".", "", -1)
		AppendRow(srv, spreadsheetId, title, price)
		time.Sleep(100)
	}

}
