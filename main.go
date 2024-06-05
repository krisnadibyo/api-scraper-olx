package main

import (
	_helper "api-scraper-olx/helper"
	"fmt"
	"time"
)

func main() {
	config := _helper.ReadConfigFile()
	searches := _helper.ReadSearchFile()
	srv, _ := _helper.SetupGsheet()

	for _, each := range searches.Search {
		fmt.Println("Fetch : " + each.SheetName)
		cars := _helper.FetchDatas(config.Url, each.Keywords)
		_helper.CreateNewSheet(srv, config.SpreadsheetID, each.SheetName)
		_helper.ClearSheet(srv, config.SpreadsheetID, each.SheetName)
		_helper.AppendRowToSheet(srv, config.SpreadsheetID, each.SheetName, _helper.AppendRowData(cars))
		_helper.AppendRowToSheet(srv, config.SpreadsheetID, each.SheetName, _helper.AppendRowFormula(cars))
		time.Sleep(1000)
	}
	// this line will cause an error because the function is not exported (lowercase first letter

}
