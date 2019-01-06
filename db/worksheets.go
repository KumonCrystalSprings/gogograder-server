package db

import (
	"strings"

	"../util"
)

// FetchWorksheetNames gets all the names of the worksheets
func FetchWorksheetNames() ([]string, error) {
	spreadsheet, err := fetchSpreadsheet(sc.Worksheets)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, sheet := range spreadsheet.Sheets {
		if strings.Contains(sheet.Properties.Title, "*") || strings.Contains(sheet.Properties.Title, "(") {
			continue
		}
		names = append(names, sheet.Properties.Title)
	}

	return names, nil
}

// WorksheetPage is the struct for the information stored in a worksheet page
type WorksheetPage struct {
	Grading []string `json:"grading"`
	Answers []string `json:"answers"`
	SideA   uint     `json:"sideA"`
	SideB   uint     `json:"sideB"`
}

// FetchWorksheetPage gets all the info about a worksheet page
func FetchWorksheetPage(ws string, page string) (ret WorksheetPage, ok bool, err error) {
	spreadsheet, err := fetchSpreadsheet(sc.Worksheets)
	if err != nil {
		return
	}

	sheet, err := spreadsheet.SheetByTitle(ws)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return
	}

	for i := 2; i < len(sheet.Rows); i++ {
		if sheet.Rows[i][2].Value == page {
			rowA := sheet.Rows[i]
			rowB := sheet.Rows[i+1]

			for j := 4; j <= 8; j++ {
				ret.Grading = append(ret.Grading, rowA[j].Value)
			}

			for j := 10; j < len(rowB); j++ {
				if rowA[j].Value != "" {
					ret.Answers = append(ret.Answers, rowA[j].Value)
					ret.SideA++
				} else if rowB[j].Value != "" {
					ret.Answers = append(ret.Answers, rowB[j].Value)
					ret.SideB++
				} else {
					break
				}
			}
			ok = true
			return
		}
	}

	return
}
