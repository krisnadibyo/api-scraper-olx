package main

import (
	_helper "api-scraper-olx/helper"
)

func main() {
	config := _helper.ReadConfigFile()

	srv, _ := _helper.SetupGsheet()
	cars := _helper.FetchDatas(config.Keywords)
	_helper.CreateNewSheet(srv, config.SpreadsheetID, config.SheetName)
	_helper.ClearSheet(srv, config.SpreadsheetID, config.SheetName)
	_helper.AppendRowToSheet(srv, config.SpreadsheetID, config.SheetName, _helper.AppendRowData(cars))
	_helper.AppendRowToSheet(srv, config.SpreadsheetID, config.SheetName, _helper.AppendRowFormula(cars)) // this line will cause an error because the function is not exported (lowercase first letter

}
