package scheduler

import (
	"github.com/my-stocks-pro/earnings-scheduler/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"github.com/my-stocks-pro/earnings-scheduler/registrator"
	"fmt"
)

type TypeReadisEarnings struct {
	ID       string
	Count    string
	Earnings string
}

type TypeScheduler struct {
	QuitOS    chan os.Signal
	QuitRPC   chan bool
	Router    *gin.Engine
	Server    *http.Server
	Config    *config.TypeConfig
	RedisData *TypeReadisEarnings
	Services  *map[string][]string
}

func ReadisEarningsNew() *TypeReadisEarnings {
	return &TypeReadisEarnings{}
}

func New() *TypeScheduler {

	services, err := registrator.GetServices()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	return &TypeScheduler{
		QuitOS:  make(chan os.Signal),
		QuitRPC: make(chan bool),
		Router:  router,
		Server: &http.Server{
			Addr:    ":8001",
			Handler: router,
		},
		Config:    config.GetConfig(),
		RedisData: ReadisEarningsNew(),
		Services:  services,
	}
}
