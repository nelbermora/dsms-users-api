package server

import (
	"github.com/gorilla/mux"
	"github.com/nelbermora/dsms-users-api/internal/controller"
	"github.com/nelbermora/dsms-users-api/internal/repository"
	"github.com/nelbermora/dsms-users-api/internal/service"
)

func setupHandlers(router *mux.Router) {
	controller := initializeController()
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{userId}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users", controller.GetUsers).Methods("OPTIONS")
	router.HandleFunc("/users/{userId}", controller.GetUsers).Methods("OPTIONS")
}

func initializeController() controller.Controller {
	userRepo, branchRepo := repository.NewRepository(), repository.NewBranchRepo()
	svc := service.NewService(userRepo, branchRepo)
	return *controller.NewController(svc)

}
