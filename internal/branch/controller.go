package branch

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

func (c *Controller) GetBranches(rw http.ResponseWriter, r *http.Request) {
	branches, err := c.svc.GetBranches()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, _ := json.Marshal(branches)
	rw.Write(data)
}

func (c *Controller) GetBranch(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchId, err := strconv.Atoi(vars["branchId"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, "Id requested error"))
		return
	}
	var branch *model.Branch
	branch, err = c.svc.GetBranch(branchId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, _ := json.Marshal(branch)
	rw.Write(data)
}

func (c *Controller) GetUsersByBranch(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	branchId, err := strconv.Atoi(vars["branchId"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, "Id requested error"))
		return
	}
	usrs, err := c.svc.GetUsersByBranch(branchId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	if usrs == nil {
		empty := []*model.User{}
		usrs = empty
	}
	data, _ := json.Marshal(usrs)
	rw.Write(data)
}

func wrapError(code int, message string) []byte {
	stError := model.NewStError(code, message)
	data, _ := json.Marshal(stError)
	return data
}
