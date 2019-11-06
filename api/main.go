package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("server up and running on port", port)

	router := httprouter.New()
	router.GET("/api", indexHandler)
	router.GET("/api/offenseplayers", getOffensePlayers)
	router.GET("/api/offenseplayers/:position", getOffensePlayers)
	router.GET("/api/kickers", getKickers)
	router.GET("/api/defenses", getDefenses)
	handler := cors.Default().Handler(router)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Println(err)
	}
}
