package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func indexHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("Leave Me Here"))
}

func getOffensePlayers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	players, err := selectOffensePlayers("")
	marshallAndSend(players, err, w)
}

func getQuarterbacks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	players, err := selectOffensePlayers("qb")
	marshallAndSend(players, err, w)
}

func getRunningBacks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	players, err := selectOffensePlayers("rb")
	marshallAndSend(players, err, w)
}

func getWideReceivers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	players, err := selectOffensePlayers("wr")
	marshallAndSend(players, err, w)
}

func getTightEnds(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	players, err := selectOffensePlayers("te")
	marshallAndSend(players, err, w)
}

func marshallAndSend(players []offensePlayer, err error, w http.ResponseWriter) {
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(players)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
