package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func main() {
	go fetchAndConvertFile()

	port := "9000"
	fmt.Println("server up and running on port", port)

	router := httprouter.New()

	router.GET("/", indexHandler)

	http.ListenAndServe(":"+port, router)
}

func fetchAndConvertFile() {
	for {
		filename := fetchFile()
		convertExcelFile(filename)
		time.Sleep(time.Minute * 60)
	}

}

func indexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("Leave Me Here"))
}
