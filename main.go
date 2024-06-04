package main

import (
	_helper "api-scraper-olx/helper"
)

func main() {
	config := _helper.ReadConfigFile()

	srv, _ := _helper.SetupGsheet()
	cars := _helper.FetchDatas(config.Keywords)
	_helper.AppendRowToSheet(srv, config.SpreadsheetID, config.SheetName, _helper.AppendRowData(cars))

}
