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
	go writeToStructAndDb()
	port := "9000"
	fmt.Println("server up and running on port", port)

	router := httprouter.New()
	router.GET("/", indexHandler)
	http.ListenAndServe(":"+port, router)
}

func fetchAndConvertFile(url string) {
	for {
		filename := fetchFile(url)
		fmt.Println("downloaded spreadsheet with filename: ", filename)
		convertExcelFile(filename)
		time.Sleep(time.Minute * 60)
	}
}

func writeToStructAndDb() {
	for {
		time.Sleep(time.Second * 10)
		go clearDb("offense_players")
		go clearDb("kickers")
		go clearDb("defenses")
		time.Sleep(time.Second * 1)
		offense := []string{"QBs", "RBs", "WRs", "TEs"}
		for _, s := range offense {
			go readToStructandDb(s)
		}
		go readToStructandDb("Ks")
		go readToStructandDb("DEFs")
		time.Sleep(time.Second * 10)
	}
}

func clearDb(s string) {
	st := fmt.Sprintf("delete from %s", s)
	sta, err := db.Prepare(st)
	if err != nil {
		panic(err)
	}
	sta.QueryRow()
	sta.Close()
}

func indexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("Leave Me Here"))
}
