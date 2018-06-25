package main

import (
	"github.com/my-stocks-pro/earnings-scheduler/scheduler"
)

func main() {

	Scheduler := scheduler.New()

	Scheduler.Run()
}
