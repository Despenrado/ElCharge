package routers

import (
	"net/http"

	controllersV1 "github.com/Despenrado/ElCharge/RestAPI/controllers/api/v1"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"gopkg.in/gorilla/mux.v1"
)

type Server struct {
	config         *utils.Config
	router         *mux.Router
	logger         *utils.Logger
	userController *controllersV1.UserController
	testController *controllersV1.TestController
}

func NewServer(config *utils.Config, ro *mux.Router, lo *utils.Logger, uc *controllersV1.UserController, tc *controllersV1.TestController) *Server {
	return &Server{
		config:         config,
		router:         ro,
		logger:         lo,
		userController: uc,
		testController: tc,
	}
}

func (s *Server) Start() error {
	s.logger.PrintInfo("Server started ...")
	if err := http.ListenAndServe(s.config.BindAddr, s.router); err != nil {
		return err
	}
	return nil
}

// SetupRouters ...
func (s *Server) SetupRouters() *mux.Router {
	v1 := "/api/v1"
	s.router.Schemes("http")
	s.router.Use(s.logger.SetRequestID)
	s.router.Use(s.logger.LogRequest)
	s.router.HandleFunc(v1, s.testController.TestAPIV1()).Methods("GET")

	user := s.router.PathPrefix(v1 + "/user").Subrouter()
	user.HandleFunc("", s.userController.CreateUser()).Methods("POST")
	return s.router
}
