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
	statController *controllersV1.StationController
	commController *controllersV1.CommentController
}

func NewServer(config *utils.Config, ro *mux.Router, lo *utils.Logger, uc *controllersV1.UserController, ac *controllersV1.AuthController, tc *controllersV1.TestController, sc *controllersV1.StationController, cc *controllersV1.CommentController) *Server {
	return &Server{
		config:         config,
		router:         ro,
		logger:         lo,
		userController: uc,
		authController: ac,
		testController: tc,
		statController: sc,
		commController: cc,
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
	s.router.HandleFunc(v1+"/login", s.authController.Login()).Methods("POST")
	s.router.HandleFunc(v1+"/logout/{id}", s.authController.Logout()).Methods("GET")

	user := s.router.PathPrefix(v1 + "/users").Subrouter()
	user.Use(s.authController.CheckToken)
	user.HandleFunc("/read", s.userController.Read()).Methods("GET")
	user.HandleFunc("/{id}", s.userController.FindByID()).Methods("GET")
	user.HandleFunc("/{id}", s.userController.DeleteByID()).Methods("DELETE")
	user.HandleFunc("/{id}", s.userController.UpdateByID()).Methods("PUT")

	stat := s.router.PathPrefix(v1 + "/stations").Subrouter()
	stat.Use(s.authController.CheckToken)
	stat.HandleFunc("", s.statController.CreateStation()).Methods("POST")
	stat.HandleFunc("/read", s.statController.Read()).Methods("GET")
	stat.HandleFunc("/{id}", s.statController.FindByID()).Methods("GET")
	stat.HandleFunc("/{id}", s.statController.DeleteByID()).Methods("DELETE")
	stat.HandleFunc("/{id}", s.statController.UpdateByID()).Methods("PUT")

	comm := stat.PathPrefix("/{sid}/comments").Subrouter()
	comm.Use(s.authController.CheckToken)
	comm.HandleFunc("/read", s.commController.Read()).Methods("GET")
	comm.HandleFunc("", s.commController.CreateComment()).Methods("POST")
	comm.HandleFunc("/{id}", s.commController.FindByID()).Methods("GET")
	comm.HandleFunc("/{id}", s.commController.DeleteByID()).Methods("DELETE")
	comm.HandleFunc("/{id}", s.commController.UpdateByID()).Methods("PUT")
	return s.router
}
