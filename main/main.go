package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	web "github.com/nelbermora/dsms-users-api/http"
)

func main() {
	router, _ := web.SetupDependencies()
	log.Println("initializing Users API")
	port := os.Getenv("PORT")
	var servePort string
	if port == "" {
		servePort = ":" + "18080"
	} else {
		servePort = ":" + port
	}

	err := http.ListenAndServe(servePort, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "GET", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowCredentials(),
	)(router))

	if err != nil {
		log.Fatal(fmt.Sprintf("could not initialize server: %s", err.Error()))
	}
}
