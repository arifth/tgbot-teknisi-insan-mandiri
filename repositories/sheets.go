package repositories

import (
	"arifthalhah/sigesit-bot/v2/models"
	"context"
	"google.golang.org/api/sheets/v4"
	"time"

	//"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/option"

	//"golang.org/x/oauth2"
	"log"
	"os"
)

type Credentials struct {
	Type                string `json:"type"`
	ProjectID           string `json:"project_id"`
	PrivateKeyID        string `json:"private_key_id"`
	PrivateKey          string `json:"private_key"`
	ClientEmail         string `json:"client_email"`
	ClientID            string `json:"client_id"`
	AuthURI             string `json:"auth_uri"`
	TokenURI            string `json:"token_uri"`
	AuthProviderCertURL string `json:"auth_provider_x509_cert_url"`
	ClientCertURL       string `json:"client_x509_cert_url"`
}

func Init() *sheets.Service {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot Print Current Directory")
	}
	path := fmt.Sprintf("%s/keys/key.json", curDir)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read service account file: %v", err)
	}
	var credential Credentials

	// Parse service account key
	err = json.Unmarshal(data, &credential)
	if err != nil {
		log.Fatalf("Unable to parse service account file: %v", err)
	}
	optionsFromFile := option.WithCredentialsFile(path)
	ctx := context.Background()

	// Create a new Sheets service
	srv, err := sheets.NewService(ctx, optionsFromFile)
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
	}
	return srv

}

func GetCellValue(srv *sheets.Service, spreadsheetId string, sheet string, rng string) (interface{}, error) {
	values, err := srv.Spreadsheets.Values.Get(spreadsheetId, fmt.Sprintf("%v!%v", sheet, rng)).Do()

	if err != nil {
		return nil, err
	}

	for _, value := range values.Values {
		fmt.Println(value)
	}
	return values.Values, nil
}

func InsertIntoSheet(srv *sheets.Service, spreadsheetId string, sheet string, rng string, data []string) (*sheets.AppendValuesResponse, error) {

	date := time.Now()
	Task := models.Job{
		Created_at: date.Format((time.ANSIC)),
	}
	for key, _ := range data {
		switch key {
		case 1:
			Task.CounterMachine = data[key]
			break
		case 2:
			break
		case len(data) - 1:
			Task.Created_by = data[len(data)-1]
		}
	}
	fmt.Println(Task)
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{},
	}
	//TODO : make not for this one, appending new array to 2 dimensional array
	newRow := make([]interface{}, len(data))
	for i, s := range data {
		newRow[i] = s
	}
	valueRange.Values = append(valueRange.Values, newRow)

	result, err := srv.Spreadsheets.Values.Append(spreadsheetId, fmt.Sprintf("%v!%v", sheet, rng), valueRange).ValueInputOption("RAW").Do()

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Helper function to get the sheet ID from the sheet name
func GetSheetID(srv *sheets.Service, spreadsheetID, sheetName string) int64 {
	resp, err := srv.Spreadsheets.Get(spreadsheetID).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve spreadsheet: %v", err)
	}

	for _, sheet := range resp.Sheets {
		if sheet.Properties.Title == sheetName {
			return sheet.Properties.SheetId
		}
	}

	log.Fatalf("Sheet not found: %s", sheetName)
	return -1 // Should not reach here
}
