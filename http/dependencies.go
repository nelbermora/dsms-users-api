package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupDependencies() (http.Handler, error) {
	router := mux.NewRouter()
	setupHandlers(router)
	return router, nil
}
