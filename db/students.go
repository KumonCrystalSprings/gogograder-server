package db

import "../util"

// FetchStudent checks if a student is in the student list and if their ID matches, and returns their record sheet ID if true
func FetchStudent(name string, password string) (string, error) {
	spreadsheet, err := fetchSpreadsheet(sc.Students)
	if err != nil {
		return "", err
	}

	sheet, err := spreadsheet.SheetByIndex(0)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return "", err
	}

	for i := 3; i < len(sheet.Rows); i++ {
		if sheet.Rows[i][1].Value == "" {
			continue
		}
		if sheet.Rows[i][1].Value == name && sheet.Rows[i][4].Value == password {
			return parseSheetURL(sheet.Rows[i][5].Value), nil
		}
	}
	return "", nil
}
