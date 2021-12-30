package company

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nelbermora/dsms-users-api/internal/model"
)

type Controller struct {
	svc Service
}

func NewController(svc Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (c *Controller) GetBranchesByCompany(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId, err := strconv.Atoi(vars["companyId"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, "Id requested error"))
		return
	}
	var branches []*model.Branch
	branches, err = c.svc.GetBranchesByCompany(companyId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, _ := json.Marshal(branches)
	rw.Write(data)
}

func (c *Controller) GetCompany(rw http.ResponseWriter, r *http.Request) {
	company, err := c.svc.GetCompany()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, _ := json.Marshal(company)
	rw.Write(data)
}

func wrapError(code int, message string) []byte {
	stError := model.NewStError(code, message)
	data, _ := json.Marshal(stError)
	return data
}
