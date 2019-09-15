package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strings"
)

func convertExcelFile(s string) {
	directory := "/tmp/"
	excelFileName := getFullFileName(directory, s)
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		if !isDisregardCategory(sheet.Name) {
			fmt.Printf("%s\n", sheet.Name)
			fileContents := []string{}
			for _, row := range sheet.Rows {
				cells := []string{}
				for _, cell := range row.Cells {
					text := cell.String() + "|"
					// fmt.Printf("%s", text)
					cells = append(cells, text)
				}
				cells[len(cells)-1] = cells[len(cells)-1][0 : len(cells[len(cells)-1])-1]
				fmt.Println(strings.Join(cells, ""))
				fileContents = append(fileContents, strings.Join(cells, ""), "\n")
			}
			writeCSVFile(directory, *sheet, []byte(strings.Join(fileContents, "")), 0755)
		}
	}
}

func getFullFileName(d string, s string) string {
	return d + s
}

func writeCSVFile(d string, s xlsx.Sheet, b []byte, p os.FileMode) {
	ioutil.WriteFile(d+s.Name, b, p)
}

func isDisregardCategory(category string) bool {
	switch strings.ToLower(category) {
	case
		"vars",
		"sheet2",
		"by post",
		"league boundaries",
		"notes":
		return true
	}
	return false
}
