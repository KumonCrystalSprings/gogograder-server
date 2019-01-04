package util

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigData struct {
	Port   int `json:"port"`
	Sheets struct {
		Students   string `json:"students"`
		Worksheets string `json:"worksheets"`
		Records    string `json:"records"`
	}
}

var Config ConfigData

func readJSON(i interface{}, path string) {
	data, err := ioutil.ReadFile(path)
	CheckErrorFatal(err, "Error in reading "+path)

	err = json.Unmarshal(data, i)
	CheckErrorFatal(err, "Error in parsing "+path)
}

func init() {
	readJSON(&Config, "config/config.json")
}
