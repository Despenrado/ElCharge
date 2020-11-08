package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/services/api"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"gopkg.in/gorilla/mux.v1"
)

type CommentController struct {
	service api.Service
}

// NewStationController constructor
func NewCommentControllerr(s api.Service) *CommentController {
	return &CommentController{
		service: s,
	}
}

func (c *CommentController) CreateComment() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		sid, ok := params["sid"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		comm := &models.Comment{}
		err := json.NewDecoder(r.Body).Decode(comm)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		comm, err = c.service.Comment().CreateComment(sid, comm)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, comm)
	})
}

func (c *CommentController) FindByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		sid, ok := params["sid"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		s, err := c.service.Comment().FindByID(sid, id)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusFound, s)
	})
}

func (c *CommentController) UpdateByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusNoContent, utils.ErrWrongRequest)
			return
		}
		sid, ok := params["sid"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		comm := &models.Comment{}
		err := json.NewDecoder(r.Body).Decode(comm)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		comm, err = c.service.Comment().UpdateByID(sid, id, comm)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, comm)
	})
}

func (c *CommentController) DeleteByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		sid, ok := params["sid"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		err := c.service.Comment().DeleteByID(sid, id)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}

func (c *CommentController) Read() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		sid, ok := vars["sid"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		params := r.URL.Query()
		skipINT, err := strconv.Atoi(params.Get("skip"))
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		limitINT, err := strconv.Atoi(params.Get("limit"))
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		stations, err := c.service.Comment().Read(sid, skipINT, limitINT)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, stations)
	})
}
