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
	authController *controllersV1.AuthController
	testController *controllersV1.TestController
}

func NewServer(config *utils.Config, ro *mux.Router, lo *utils.Logger, uc *controllersV1.UserController, ac *controllersV1.AuthController, tc *controllersV1.TestController) *Server {
	return &Server{
		config:         config,
		router:         ro,
		logger:         lo,
		userController: uc,
		authController: ac,
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
	s.router.Use(s.logger.SetRequestID) // middleware
	s.router.Use(s.logger.LogRequest)   // middleware
	s.router.HandleFunc(v1, s.testController.TestAPIV1()).Methods("GET")
	s.router.HandleFunc(v1+"/users", s.authController.CreateUser()).Methods("POST")
	s.router.HandleFunc(v1+"/login", s.authController.Login()).Methods("GET")
	s.router.HandleFunc(v1+"/logout/{id}", s.authController.Logout()).Methods("GET")

	user := s.router.PathPrefix(v1 + "/users").Subrouter()
	user.Use(s.authController.CheckToken)
	user.HandleFunc("/{id}", s.userController.FindByID()).Methods("GET")
	user.HandleFunc("/{id}", s.userController.DeleteByID()).Methods("DELETE")
	user.HandleFunc("/{id}", s.userController.UpdateByID()).Methods("PUT")
	return s.router
}
