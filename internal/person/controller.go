package person

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

func (c *Controller) GetUsers(rw http.ResponseWriter, r *http.Request) {
	usrs, err := c.svc.GetUsers()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, _ := json.Marshal(usrs)
	rw.Write(data)
}

func (c *Controller) GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, "Id requested error"))
		return
	}
	usr, err := c.svc.GetUser(userId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	if usr == nil {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write(wrapError(http.StatusNotFound, "User not found"))
		return
	}
	data, _ := json.Marshal(usr)
	rw.Write(data)
}

func (c *Controller) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var usr model.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, "Bad Request"))
		return
	}
	err = usr.Validate()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, err.Error()))
		return
	}
	var updated *model.User
	updated, err = c.svc.UpdateUser(usr)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, _ := json.Marshal(updated)
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

func (c *Controller) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, "Id requested error"))
		return
	}
	err = c.svc.DeleteUser(userId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	rw.WriteHeader(http.StatusNoContent)
}

func (c *Controller) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var usr model.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, "Bad Request"))
		return
	}
	err = usr.Validate()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(wrapError(http.StatusBadRequest, err.Error()))
		return
	}
	var created *model.User
	created, err = c.svc.CreateUser(usr)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(wrapError(http.StatusInternalServerError, err.Error()))
		return
	}
	data, _ := json.Marshal(created)
	rw.WriteHeader(http.StatusCreated)
	rw.Write(data)
}

func wrapError(code int, message string) []byte {
	stError := model.NewStError(code, message)
	data, _ := json.Marshal(stError)
	return data
}
