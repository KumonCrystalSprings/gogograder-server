package db

import (
	"fmt"
	"time"

	"../util"

	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

type StudentActivity struct {
	Name         string    `json:"name"`
	Sheet        string    `json:"sheet"`
	Subject      string    `json:"subject"`
	Date         time.Time `json:"date"`
	Worksheet    string    `json:"worksheet"`
	Page         uint      `json:"page"`
	Score        uint      `json:"score"`
	Time         uint      `json:"time"`
	ReportedTime uint      `json:"reportedTime"`
}

func WriteStudentActivity(a *StudentActivity) error {
	recordSheet, err := fetchSpreadsheet(sc.Records)
	if err != nil {
		return err
	}

	recordTab, err := recordSheet.SheetByTitle("RT " + a.Subject)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return err
	}

	studentSheet, err := fetchSpreadsheet(a.Sheet)
	if err != nil {
		return err
	}

	studentTab, err := studentSheet.SheetByTitle("RT " + a.Subject)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return err
	}

	err1 := addActivity(a, recordTab)
	err2 := addActivity(a, studentTab)

	if err1 != nil {
		return err1
	}

	return err2
}

func addActivity(a *StudentActivity, tab *spreadsheet.Sheet) error {
	index := len(tab.Rows)

	tab.Update(index, 0, a.Name)
	tab.Update(index, 1, a.Date.Format("1/2/2006"))
	tab.Update(index, 3, a.Date.Format("3:04 PM"))
	tab.Update(index, 4, a.Worksheet)
	tab.Update(index, 5, fmt.Sprint(a.Page))
	tab.Update(index, 6, fmt.Sprint(a.Score))
	tab.Update(index, 8, fmt.Sprint(a.Time))
	tab.Update(index, 10, fmt.Sprint(a.ReportedTime))

	return tab.Synchronize()
}
