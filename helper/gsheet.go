package helper

import (
	"context"
	"encoding/json"
//	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Token struct {
	PrivateKey   string `json:"private_key"`
	PrivateKeyID string `json:"private_key_id"`
	Email        string `json:"client_email"`
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

func AppendRow(srv *sheets.Service, spreadsheetId string) {

	values := &sheets.ValueRange{
		Values: [][]interface{}{{
		"Timmy",
		"Malaysia",
		}},
	}

	_, err := srv.Spreadsheets.Values.Append(spreadsheetId, "Sheet1!A:F", values).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		log.Fatalf("Unable to insert data to sheet: %v", err)
	}
}

func maingo() {

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
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "13QBSQflOCao6HYJucZ02YnB8Uq5Q9l4cVEN0ab1VOXw"
	AppendRow(srv, spreadsheetId)
	
}