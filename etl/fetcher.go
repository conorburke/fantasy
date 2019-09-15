package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func fetchFile() string {
	url := "http://walterfootball.com/fantasy2019rankingsexcel.xlsx"
	fmt.Println("url", url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	re := regexp.MustCompile(`([^\/]+$)`)
	filename := re.FindString(url)

	file, err := os.Create(fmt.Sprintf("/tmp/%s", filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)

	return filename
}
