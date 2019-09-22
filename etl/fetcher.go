package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func fetchFile(url string) string {
	fmt.Println("url", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	filename := getSpreadsheetName(url)

	file, err := os.Create(fmt.Sprintf("/tmp/%s", filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	return filename
}

func getSpreadsheetName(url string) string {
	re := regexp.MustCompile(`([^\/]+$)`)
	return re.FindString(url)
}
