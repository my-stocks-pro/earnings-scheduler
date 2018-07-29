package main

import (
	"github.com/my-stocks-pro/earnings-scheduler/scheduler"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	Scheduler := scheduler.New()

	go Scheduler.Run()

	router.Run(":8002")
}
