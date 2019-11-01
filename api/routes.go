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

func getOffensePlayers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	players, err := selectOffensePlayers(ps.ByName("position"))
	marshallAndSendOffense(players, err, w)
}

// func getOffensePlayers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	players, err := selectOffensePlayers("")
// 	marshallAndSendOffense(players, err, w)
// }

// func getQuarterbacks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	players, err := selectOffensePlayers("qb")
// 	marshallAndSendOffense(players, err, w)
// }

// func getRunningBacks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	players, err := selectOffensePlayers("rb")
// 	marshallAndSendOffense(players, err, w)
// }

// func getWideReceivers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	players, err := selectOffensePlayers("wr")
// 	marshallAndSendOffense(players, err, w)
// }

// func getTightEnds(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	players, err := selectOffensePlayers("te")
// 	marshallAndSendOffense(players, err, w)
// }

func getKickers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	kickers, err := selectKickers()
	marshallAndSendKicker(kickers, err, w)
}

func getDefenses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defenses, err := selectDefenses()
	marshallAndSendDefense(defenses, err, w)
}

func marshallAndSendOffense(players []offensePlayer, err error, w http.ResponseWriter) {
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

func marshallAndSendKicker(kickers []kicker, err error, w http.ResponseWriter) {
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(kickers)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func marshallAndSendDefense(defenses []defense, err error, w http.ResponseWriter) {
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(defenses)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(js)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}