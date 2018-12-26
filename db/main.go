package db

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"

	"../util"
)

// SheetsConfig represents the sheets.json config file
type SheetsConfig struct {
	Students   string `json:"students"`
	Worksheets string `json:"worksheets"`
	Records    string `json:"records"`
}

var service *spreadsheet.Service
var sc SheetsConfig

func init() {
	// initialize google sheets service
	data, err := ioutil.ReadFile("config/client_secret.json")
	util.CheckErrorFatal(err, "Error in reading client secret file")

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	util.CheckErrorFatal(err, "Error in verifying client secret")

	client := conf.Client(context.TODO())
	service = spreadsheet.NewServiceWithClient(client)

	// get sheet ID's
	data, err = ioutil.ReadFile("config/sheets.json")
	util.CheckErrorFatal(err, "Error in reading Google Sheet config file")

	err = json.Unmarshal(data, &sc)
	util.CheckErrorFatal(err, "Error in parsing Google Sheet config file")
}

func fetchSpreadsheet(id string) (spreadsheet.Spreadsheet, error) {
	spreadsheet, err := service.FetchSpreadsheet(id)
	util.CheckError(err, "Error in fetching spreadsheet: https://docs.google.com/spreadsheets/d/"+id+"/edit")
	return spreadsheet, err
}
