package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func main() {
	//may make this a command line arg in future
	url := "http://walterfootball.com/fantasy2019rankingsexcel.xlsx"
	go fetchAndConvertFile(url)

	port := "9000"
	fmt.Println("server up and running on port", port)

	router := httprouter.New()

	router.GET("/", indexHandler)

	http.ListenAndServe(":"+port, router)
}

func fetchAndConvertFile(url string) {
	for {
		filename := fetchFile(url)
		fmt.Println("filename", filename)
		convertExcelFile(filename)
		time.Sleep(time.Minute * 60)
	}

}

func indexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("Leave Me Here"))
}
