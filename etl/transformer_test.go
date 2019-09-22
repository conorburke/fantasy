package main

import "testing"

func TestGetFullFileName(t *testing.T) {
	directory := "/tmp/"
	fname := "banjokazooie.xlsx"
	fpath := getFullFileName(directory, fname)

	t.Log(fpath)

	if fpath != "/tmp/banjokazooie.xlsx" {
		t.Error("File naming is not correct")
	} else {
		t.Log("File name is correct")
	}
}
