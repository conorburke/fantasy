package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("server up and running on port", port)

	router := httprouter.New()
	router.GET("/", indexHandler)
	router.GET("/offenseplayers", getOffensePlayers)
	router.GET("/quarterbacks", getQuarterbacks)
	router.GET("/runningbacks", getRunningBacks)
	router.GET("/widereceivers", getWideReceivers)
	router.GET("/tightends", getTightEnds)
	http.ListenAndServe(":"+port, router)
}
