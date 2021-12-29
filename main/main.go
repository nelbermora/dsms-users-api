package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nelbermora/dsms-users-api/internal/server"
)

func main() {
	router, _ := server.SetupDependencies()
	log.Println("initializing Users API")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(fmt.Sprintf("could not initialize server: %s", err.Error()))
	}
}
