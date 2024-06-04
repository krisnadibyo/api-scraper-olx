package main

import (
	_helper "api-scraper-olx/helper"
)

func main() {
	spreadsheetId := "13QBSQflOCao6HYJucZ02YnB8Uq5Q9l4cVEN0ab1VOXw"
	srv, _ := _helper.SetupGsheet()
	cars := _helper.FetchData()
	_helper.AppendRowToSheet(srv, spreadsheetId, _helper.AppendRowData(cars))

}
