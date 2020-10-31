package main

import (
	"flag"
	"log"

	controllersV1 "github.com/Despenrado/ElCharge/RestAPI/controllers/api/v1"
	"github.com/Despenrado/ElCharge/RestAPI/routers"
	servicesV1 "github.com/Despenrado/ElCharge/RestAPI/services/api/v1"
	mongostorage "github.com/Despenrado/ElCharge/RestAPI/storage/mongostore"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorilla/mux.v1"
)

// TODO: Finish up the **base structure** o project
// - [ ] main
// - [X] models
// - [X] controllers
// - [X] storage
// - [ ] services
// - [X] routers
// - [ ] configs
// - [X] dockerfile

var (
	configPath string
)

// Config ...
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.yaml", "path to config file")
}

func createServerDependencies(path string) *routers.Server {
	config, err := utils.NewConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	ur := mongostorage.NewUserRepository(nil)
	st := mongostorage.NewStorage(ur)
	us := servicesV1.NewUserService(st)
	se := servicesV1.NewService(us)
	uc := controllersV1.NewUserController(se)
	tc := controllersV1.NewTestController(se)
	ro := mux.NewRouter()
	lo := &utils.Logger{*logrus.New()}
	s := routers.NewServer(config, ro, lo, uc, tc)
	return s
}

func main() {
	flag.Parse()
	s := createServerDependencies(configPath)
	s.SetupRouters()
	s.Start()
}
