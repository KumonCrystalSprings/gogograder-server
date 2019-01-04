package db

import (
	"fmt"
	"time"

	"../util"
)

type JSONDate time.Time

func (d JSONDate) MarshalJSON() ([]byte, error) {
	dateString := time.Time(d).Format("1/2/2006")
	return []byte(dateString), nil
}

func (d JSONDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	t, err := time.Parse("1/2/2006", s)
	d = JSONDate(t)
	return err
}

func (d JSONDate) String() string {
	return time.Time(d).Format("1/2/2006")
}

type StudentActivity struct {
	Name      string   `json:"name"`
	Date      JSONDate `json:"date"`
	Worksheet string   `json:"worksheet"`
	Page      uint     `json:"page"`
	Score     uint     `json:"score"`
	Time      uint     `json:"time"`
}

func WriteStudentActivity(a *StudentActivity) error {
	spreadsheet, err := fetchSpreadsheet(sc.Records)
	if err != nil {
		return err
	}

	sheet, err := spreadsheet.SheetByIndex(0)
	if util.CheckError(err, "Error in fetching students spreadsheet") {
		return err
	}

	index := len(sheet.Rows)
	sheet.Update(index, 0, a.Name)
	sheet.Update(index, 1, fmt.Sprint(a.Date))
	sheet.Update(index, 2, a.Worksheet)
	sheet.Update(index, 3, fmt.Sprint(a.Page))
	sheet.Update(index, 4, fmt.Sprint(a.Score))
	sheet.Update(index, 5, fmt.Sprint(a.Time))

	return sheet.Synchronize()
}
