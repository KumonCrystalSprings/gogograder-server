package db

import "../util"

// CheckStudent checks if a student is in the student list and if their ID matches
func CheckStudent(name string, id string) (bool, error) {
	spreadsheet, err := service.FetchSpreadsheet(sc.Students)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return false, err
	}

	sheet, err := spreadsheet.SheetByIndex(0)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return false, err
	}

	for _, r := range sheet.Rows {
		if r[0].Value == name && r[1].Value == id {
			return true, nil
		}
	}
	return false, nil
}
