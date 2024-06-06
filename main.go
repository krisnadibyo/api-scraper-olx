package main

import (
	_helper "api-scraper-olx/helper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// func main() {
// 	config := _helper.ReadConfigFile()
// 	searches := _helper.ReadSearchFile()
// 	srv, _ := _helper.SetupGsheet()

// 	for _, each := range searches.Search {
// 		fmt.Println("Fetch : " + each.SheetName)
// 		cars := _helper.FetchDatas(config.Url, each.Keywords)
// 		_helper.CreateNewSheet(srv, config.SpreadsheetID, each.SheetName)
// 		_helper.ClearSheet(srv, config.SpreadsheetID, each.SheetName)
// 		_helper.AppendRowToSheet(srv, config.SpreadsheetID, each.SheetName, _helper.AppendRowData(cars))
// 		_helper.AppendRowToSheet(srv, config.SpreadsheetID, each.SheetName, _helper.AppendRowFormula(cars))
// 		time.Sleep(1000)
// 	}
// 	// this line will cause an error because the function is not exported (lowercase first letter

// }

type SearchRequest struct {
	Keywords  string `json:"keywords"`
	SheetName string `json:"sheet_name"`
}

func DoScrape(keywords string, sheetName string) {
	config := _helper.ReadConfigFile()
	srv, _ := _helper.SetupGsheet()
	searches := strings.Split(keywords, ",")
	cars := _helper.FetchDatas(config.Url, searches)
	_helper.CreateNewSheet(srv, config.SpreadsheetID, sheetName)
	_helper.ClearSheet(srv, config.SpreadsheetID, sheetName)
	_helper.AppendRowToSheet(srv, config.SpreadsheetID, sheetName, _helper.AppendRowData(cars))
	_helper.AppendRowToSheet(srv, config.SpreadsheetID, sheetName, _helper.AppendRowFormula(cars))

}

func scrape(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Endpoint Hit: scrape")
	if r.Method == "POST" {
		reqBody, _ := ioutil.ReadAll(r.Body)
		var searchReq SearchRequest
		json.Unmarshal(reqBody, &searchReq)
		fmt.Println(searchReq)
		DoScrape(searchReq.Keywords, searchReq.SheetName)

		result := map[string]string{"status": "success", "message": "Scraping data success"}
		response, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

}

func main() {
	http.HandleFunc("/scrape", scrape)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
