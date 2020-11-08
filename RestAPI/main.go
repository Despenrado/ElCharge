package main

import (
	"flag"
	"log"

	controllersV1 "github.com/Despenrado/ElCharge/RestAPI/controllers/api/v1"
	"github.com/Despenrado/ElCharge/RestAPI/routers"
	servicesV1 "github.com/Despenrado/ElCharge/RestAPI/services/api/v1"
	mongostorage "github.com/Despenrado/ElCharge/RestAPI/storage/mongostore"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorilla/mux.v1"
)

// TODO: Finish up the **base structure** o project
// - [ ] main
// - [X] models
// - [ ] controllers
// - [X] storage
// - [ ] services
// - [X] routers
// - [X] configs
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
	db, err := mongostorage.ConnectToDB(config.DatabaseURL, config.DbName)
	if err != nil {
		log.Fatal(err)
	}
	ur := mongostorage.NewUserRepository(mongostorage.ConfigureRepository(db, config.DbUserCollection))
	sr := mongostorage.NewStationRepository(mongostorage.ConfigureRepository(db, config.DbStationCollection))
	cr := mongostorage.NewCommentRepository(sr)
	st := mongostorage.NewStorage(ur, sr, cr)
	us := servicesV1.NewUserService(st)
	ss := servicesV1.NewStationService(st)
	cs := servicesV1.NewCommentService(st)
	se := servicesV1.NewService(us, ss, cs)
	uc := controllersV1.NewUserController(se)
	sc := controllersV1.NewStationController(se)
	cc := controllersV1.NewCommentControllerr(se)
	rcli := redis.NewClient(&redis.Options{
		Addr: config.RedisDB,
	})
	_, err = rcli.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	ac := controllersV1.NewAuthController(se, rcli)
	ac.SetJWTKey(config.JWTKey)
	tc := controllersV1.NewTestController(se)
	ro := mux.NewRouter()
	lo := &utils.Logger{*logrus.New()}
	s := routers.NewServer(config, ro, lo, uc, ac, tc, sc, cc)
	return s
}

func main() {
	flag.Parse()
	s := createServerDependencies(configPath)
	s.SetupRouters()
	s.Start()
}
