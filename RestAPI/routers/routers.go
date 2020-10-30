package routers

import (
	apiv1 "github.com/Despenrado/ElCharge/RestAPI/controllers/api/v1"
	"gopkg.in/gorilla/mux.v1"
)

// SetupRouters ...
func SetupRouters() *mux.Router {
	r := mux.NewRouter()
	v1 := "/api/v1"
	r.HandleFunc(v1, apiv1.TestAPIV1()).Methods("GET")
	return nil
}
