package scheduler

import (
	"github.com/my-stocks-pro/earnings-scheduler/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
}

func ReadisEarningsNew() *TypeReadisEarnings {
	return &TypeReadisEarnings{}
}

func New() *TypeScheduler {
	router := gin.Default()
	return &TypeScheduler{
		QuitOS:  make(chan os.Signal),
		QuitRPC: make(chan bool),
		Router:  router,
		Server: &http.Server{
			Addr:    ":8002",
			Handler: router,
		},
		Config:    config.GetConfig(),
		RedisData: ReadisEarningsNew(),
	}
}
