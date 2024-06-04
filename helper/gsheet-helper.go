package helper

import (
	"context"
	"strings"

	//	"fmt"
	"log"

	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Token struct {
	PrivateKey   string `json:"private_key"`
	PrivateKeyID string `json:"spreadsheetId"`
	Email        string `json:"client_email"`
}

func SetupGsheet() (*sheets.Service, error) {
	token := ReadTokenFile()

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

func AppendRowData(data []Item) *sheets.ValueRange {
	var vr sheets.ValueRange
	var v [][]interface{}

	for _, each := range data {
		price := each.Price.Value.Display
		price = price[3:]
		price = strings.Replace(price, ".", "", -1)
		v = append(v, []interface{}{each.Title, price})
	}

	vr.Values = v
	return &vr
}

func AppendRowToSheet(srv *sheets.Service, spreadsheetId string, sheetName string, rowData *sheets.ValueRange) {
	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, sheetName+"!A:F", rowData).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		log.Fatalf("Unable to insert data to sheet: %v", err)
	}
}

func CreateNewSheet(srv *sheets.Service, spreadsheetId string, sheetName string) {
	//check if sheetname already exist
	// Retrieve the spreadsheet to get the list of sheets.
	spreadsheet, err := srv.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve spreadsheet: %v", err)
	}

	// Check if the sheet already exists.
	sheetExists := false
	for _, sheet := range spreadsheet.Sheets {
		if sheet.Properties.Title == sheetName {
			sheetExists = true
			break
		}
	}

	if sheetExists {
		log.Printf("Sheet %s already exists", sheetName)
		return
	}

	rb := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			&sheets.Request{
				AddSheet: &sheets.AddSheetRequest{
					Properties: &sheets.SheetProperties{
						Title: sheetName,
					},
				},
			},
		},
	}

	_, er := srv.Spreadsheets.BatchUpdate(spreadsheetId, rb).Do()
	if er != nil {
		log.Fatalf("Unable to create new sheet: %v", err)
	}
}
