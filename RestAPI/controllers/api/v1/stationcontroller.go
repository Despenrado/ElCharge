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

type StationController struct {
	service api.Service
}

// NewStationController constructor
func NewStationController(s api.Service) *StationController {
	return &StationController{
		service: s,
	}
}

func (c *StationController) CreateStation() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := &models.Station{}
		err := json.NewDecoder(r.Body).Decode(s)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		s, err = c.service.Station().CreateStation(s)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		utils.Respond(w, r, http.StatusCreated, s)
	})
}

func (c *StationController) FindByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		s, err := c.service.Station().FindByID(id)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusFound, s)
	})
}

func (c *StationController) UpdateByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		query := r.URL.Query()
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusNoContent, utils.ErrWrongRequest)
			return
		}
		ownid := query.Get("ownid")
		s := &models.Station{}
		err := json.NewDecoder(r.Body).Decode(s)
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		s, err = c.service.Station().UpdateByID(id, s, ownid)
		if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, s)
	})
}

func (c *StationController) DeleteByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		query := r.URL.Query()
		id, ok := params["id"]
		if !ok {
			utils.Error(w, r, http.StatusBadRequest, utils.ErrWrongRequest)
			return
		}
		ownid := query.Get("ownid")
		err := c.service.Station().DeleteByID(id, ownid)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, nil)
	})
}

func (c *StationController) Read() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		name := params.Get("name")
		if name != "" {
			stations, err := c.service.Station().FindByName(name)
			if err != nil {
				utils.Error(w, r, http.StatusNoContent, err)
				return
			}
			utils.Respond(w, r, http.StatusOK, stations)
			return
		}
		descr := params.Get("descr")
		if name != "" {
			stations, err := c.service.Station().FindByDescription(descr)
			if err != nil {
				utils.Error(w, r, http.StatusNoContent, err)
				return
			}
			utils.Respond(w, r, http.StatusOK, stations)
			return
		}
		latitude, err := strconv.ParseFloat(params.Get("lat"), 64)
		if err == nil {
			longitude, err := strconv.ParseFloat(params.Get("lng"), 64)
			if err == nil {
				// if we want to find station in radius around the coordinates
				distance, err := strconv.Atoi(params.Get("dist"))
				if err == nil && distance != 0 {
					stations, err := c.service.Station().FindInRadius(latitude, longitude, distance, skipINT, limitINT)
					if err != nil {
						utils.Error(w, r, http.StatusNoContent, err)
						return
					}
					utils.Respond(w, r, http.StatusOK, stations)
					return
				}
				// if we want get station in coordinates
				station, err := c.service.Station().FindByLocation(latitude, longitude)
				if err != nil {
					utils.Error(w, r, http.StatusNoContent, err)
					return
				}
				utils.Respond(w, r, http.StatusOK, station)
				return
			}
		}
		stations, err := c.service.Station().Read(skipINT, limitINT)
		if err != nil {
			utils.Error(w, r, http.StatusNoContent, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, stations)
	})
}
