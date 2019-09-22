package main

import "testing"

func TestGetSpreadsheetName(t *testing.T) {
	url := "https://somewebsite.com/banjokazooie.xlsx"
	filename := getSpreadsheetName(url)

	t.Log(filename)

	if filename != "banjokazooie.xlsx" {
		t.Error("Spreadsheet naming is not correct")
	} else {
		t.Log("Spreadsheet name is correct")
	}
}
