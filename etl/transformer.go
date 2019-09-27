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
		if isDesiredCategory(sheet.Name) {
			fmt.Printf("sheet name sent to %s : %s\n", directory, sheet.Name)
			fileContents := []string{}
			//vbd custom is the final column we want
			vbd := 1
			for i, row := range sheet.Rows {
				cells := []string{}
				for _, cell := range row.Cells {
					text := cell.String() + "|"
					cells = append(cells, text)
				}

				// find the index of the vbd custom column and cut off the slices at that index
				if i == 0 {
					for j, v := range cells {
						if strings.ToLower(v) == "vbd custom|" {
							vbd += j
						}
					}
				}
				cells = cells[0:vbd]

				//remove the trailing pipe
				cells[len(cells)-1] = cells[len(cells)-1][0 : len(cells[len(cells)-1])-1]
				fileContents = append(fileContents, strings.Join(cells, ""), "\n")
			}
			// fmt.Println(fileContents)
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

func isDesiredCategory(category string) bool {
	switch strings.ToLower(category) {
	case
		"qbs",
		"rbs",
		"wrs",
		"tes",
		"ks",
		"defs":
		return true
	}
	return false
}
