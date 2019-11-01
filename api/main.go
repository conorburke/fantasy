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
	router.GET("/offenseplayers/:position", getOffensePlayers)
	router.GET("/kickers", getKickers)
	router.GET("/defenses", getDefenses)
	http.ListenAndServe(":"+port, router)
}
