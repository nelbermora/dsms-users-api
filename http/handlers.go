package http

import (
	"github.com/gorilla/mux"
	"github.com/nelbermora/dsms-users-api/internal/branch"
	"github.com/nelbermora/dsms-users-api/internal/company"
	"github.com/nelbermora/dsms-users-api/internal/person"
)

func setupHandlers(router *mux.Router) {
	personCtl := initializePerson()
	router.HandleFunc("/persons", personCtl.GetUsers).Methods("GET")
	router.HandleFunc("/persons/{userId}", personCtl.GetUser).Methods("GET")
	router.HandleFunc("/persons", personCtl.CreateUser).Methods("POST")
	router.HandleFunc("/persons", personCtl.UpdateUser).Methods("PUT")
	router.HandleFunc("/persons/{userId}", personCtl.DeleteUser).Methods("DELETE")
	router.HandleFunc("/persons", personCtl.GetUsers).Methods("OPTIONS")
	router.HandleFunc("/persons/{userId}", personCtl.GetUsers).Methods("OPTIONS")
	branchCtl := initializeBranch()
	router.HandleFunc("/branches/{branchId}", branchCtl.GetBranch).Methods("GET")
	router.HandleFunc("/branches/{branchId}/persons", branchCtl.GetUsersByBranch).Methods("GET")
	companyCtl := initializeCompany()
	router.HandleFunc("/company", companyCtl.GetCompany).Methods("GET")
	router.HandleFunc("/company/{companyId}/branches", companyCtl.GetBranchesByCompany).Methods("GET")
}

func initializePerson() person.Controller {
	userRepo := person.NewRepository()
	svc := person.NewService(userRepo)
	return *person.NewController(svc)
}

func initializeBranch() branch.Controller {
	repo := branch.NewBranchRepo()
	userRepo := person.NewRepository()
	svc := branch.NewService(repo, userRepo)
	return *branch.NewController(svc)
}

func initializeCompany() company.Controller {
	branchRepo := branch.NewBranchRepo()
	repo := company.NewCompanyRepo()
	svc := company.NewService(repo, branchRepo)
	return *company.NewController(svc)
}
