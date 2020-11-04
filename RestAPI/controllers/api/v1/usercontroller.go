package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/services/api"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"gopkg.in/gorilla/mux.v1"
)

type UserController struct {
	service api.Service
}

// NewUserController constructor
func NewUserController(s api.Service) *UserController {
	return &UserController{
		service: s,
	}
}

func (c *UserController) CreateUser() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u := &models.User{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err = c.service.User().CreateUser(u)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, u)
	})
}

func (c *UserController) FindByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		u, err := c.service.User().FindByID(id)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusFound, u)
	})
}

func (c *UserController) UpdateByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusNoContent, utils.ErrWrongRequest)
			return
		}
		u := &models.User{}
		err := json.NewDecoder(r.Body).Decode(u)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err = c.service.User().UpdateByID(id, u)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, u)
	})
}

func (c *UserController) DeleteByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		err := c.service.User().DeleteByID(id)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}
