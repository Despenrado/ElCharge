package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/services/api"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
)

type UserController struct {
	service api.Service
}

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
		utils.Respond(w, r, http.StatusOK, u)
	})
}
