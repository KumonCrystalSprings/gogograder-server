package db

import (
	"context"
	"io/ioutil"
	"strings"

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

	sc = util.Config.Sheets

	sc.Students = parseSheetURL(sc.Students)
	sc.Worksheets = parseSheetURL(sc.Worksheets)
	sc.Records = parseSheetURL(sc.Records)
}

func fetchSpreadsheet(id string) (spreadsheet.Spreadsheet, error) {
	spreadsheet, err := service.FetchSpreadsheet(id)
	util.CheckError(err, "Error in fetching spreadsheet: https://docs.google.com/spreadsheets/d/"+id+"/edit")
	return spreadsheet, err
}

func parseSheetURL(url string) (id string) {
	start := strings.LastIndex(url, "/d/") + len("/d/")
	end := strings.LastIndex(url, "/edit")
	if start == -1 || end == -1 {
		return
	}
	id = url[start:end]
	return
}
