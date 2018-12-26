package db

import "../util"

// CheckStudent checks if a student is in the student list and if their ID matches
func CheckStudent(name string, id string) (bool, error) {
	spreadsheet, err := fetchSpreadsheet(sc.Students)
	if err != nil {
		return false, err
	}

	sheet, err := spreadsheet.SheetByIndex(0)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return false, err
	}

	for i := 1; i < len(sheet.Rows); i++ {
		if sheet.Rows[i][0].Value == "" {
			continue
		}
		if sheet.Rows[i][0].Value == name && sheet.Rows[i][1].Value == id {
			return true, nil
		}
	}
	return false, nil
}
