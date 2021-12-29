package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nelbermora/dsms-users-api/internal/server"
)

func main() {
	router, _ := server.SetupDependencies()
	log.Println("initializing Users API")
	port := os.Getenv("PORT")
	var servePort string
	if port == "" {
		servePort = ":" + "18080"
	} else {
		servePort = ":" + port
	}
	err := http.ListenAndServe(servePort, router)
	if err != nil {
		log.Fatal(fmt.Sprintf("could not initialize server: %s", err.Error()))
	}
}
