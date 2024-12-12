package repositories

import (
	"context"
	"google.golang.org/api/sheets/v4"

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

func GetCellValue(srv *sheets.Service, spreadsheetId string) (interface{}, error) {
	values, err := srv.Spreadsheets.Values.Get(spreadsheetId, "Sheet1!A2:E7").Do()

	if err != nil {
		return nil, err
	}

	for _, value := range values.Values {
		fmt.Println(value)
	}
	return values.Values, nil
}

func InsertIntoSheet(srv *sheets.Service, spreadsheetId string, data []string) (*sheets.AppendValuesResponse, error) {

	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{},
	}
	//TODO : make not for this one, appending new array to 2 dimensional array
	newRow := make([]interface{}, len(data))
	for i, s := range data {
		newRow[i] = s
	}
	valueRange.Values = append(valueRange.Values, newRow)

	result, err := srv.Spreadsheets.Values.Append(spreadsheetId, "Sheet1!A1:M1", valueRange).ValueInputOption("RAW").Do()

	if err != nil {
		return nil, err
	}

	//for _, value := range result.TableRange {
	//	fmt.Println(value)
	//}

	return result, nil
}
