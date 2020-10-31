package v1

import (
	"net/http"

	"github.com/Despenrado/ElCharge/RestAPI/services/api"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
)

type TestController struct {
	service api.Service
}

func NewTestController(s api.Service) *TestController {
	return &TestController{
		service: s,
	}
}

// TestAPIV1 ...
func (c *TestController) TestAPIV1() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Respond(w, r, http.StatusOK, "test")
	})
}
