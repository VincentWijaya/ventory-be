package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/vincentwijaya/ventory-be/internal/app/handler"

	"github.com/go-chi/chi"
	userRepo "github.com/vincentwijaya/ventory-be/internal/app/repo/user"
	userUC "github.com/vincentwijaya/ventory-be/internal/app/usecase/user"
	"github.com/vincentwijaya/ventory-be/pkg/database"
	"github.com/vincentwijaya/ventory-be/pkg/log"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Server   ServerConfig
	Log      LogConfig
	Database database.Config
}

type ServerConfig struct {
	Port        string
	Environment string
	JwtSecret   string
}

type LogConfig struct {
	LogPath string
	Level   string
}

const fileLocation = "/etc/ventory/"
const devFileLocation = "files/etc/ventory/"
const fileName = "ventory.%s.yaml"
const devFileName = "ventory.yaml"

const infoFile = "ventory.info.log"
const errorFile = "ventory.error.log"

func main() {
	//Read config
	var config Config
	location, fileName := getConfigLocation()
	err := gcfg.ReadFileInto(&config, location+fileName)
	if err != nil {
		log.Error("Failed to start service:", err)
		return
	}

	logConfig := log.LogConfig{
		StdoutFile: config.Log.LogPath + infoFile,
		StderrFile: config.Log.LogPath + errorFile,
		Level:      config.Log.Level,
	}
	log.InitLogger(config.Server.Environment, logConfig, []string{})

	masterDB, err := database.Connect(config.Database)
	if err != nil {
		log.Error("Failed to connect master database:", err)
		return
	}

	// Repository
	user := userRepo.New(masterDB)

	// Usecase
	userUsecase := userUC.New(user, config.Server.JwtSecret)

	// Hanlder
	httpHandler := handler.New(userUsecase)

	fmt.Printf("%+v", masterDB)

	httpRouter := chi.NewRouter()
	checker := systemCheck{
		pinger: map[string]Tester{
			"MasterDB": masterDB,
		},
	}

	httpRouter.Get("/ping", checker.ping)
	httpRouter.Get("/health", checker.health)

	httpRouter.Route("/v1", func(r chi.Router) {
		r.Post("/login", httpHandler.Login)
	})

	log.Infof("Service Started on:%v", config.Server.Port)
	err = http.ListenAndServe(config.Server.Port, httpRouter)
	if err != nil {
		log.Info("Failed serving Chi Dispatcher:", err)
		return
	}
	log.Info("Serving Chi Dispatcher on port:", config.Server.Port)
}

//-----------[ Pinger ]-----------------

type Tester interface {
	Ping() error
}

type systemCheck struct {
	pinger map[string]Tester
}

func (sys *systemCheck) ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func (sys *systemCheck) health(w http.ResponseWriter, r *http.Request) {
	var str string
	for k, v := range sys.pinger {
		start := time.Now()
		status := "Success"
		message := "successful"
		if err := v.Ping(); err != nil {
			status = "Error"
			message = err.Error()
		}
		duration := time.Now().Sub(start).Nanoseconds()
		str = fmt.Sprintf("%s%s | %s | %s | %dms\n", str, k, status, message, duration)
	}
	_, _ = w.Write([]byte(str))
}

func getConfigLocation() (string, string) {
	env := os.Getenv("PRAKERJA_ENV")
	location := devFileLocation
	name := devFileName
	if env == "staging" || env == "production" || env == "development" {
		location = fileLocation
		name = fmt.Sprintf(fileName, env)
	}
	return location, name
}
