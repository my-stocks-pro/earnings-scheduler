package scheduler

import "github.com/my-stocks-pro/earnings-scheduler/config"

type TypeReadisEarnings struct {
	ID       string
	Count    string
	Earnings string
}

type TypeScheduler struct {
	Config    *config.TypeConfig
	RedisData *TypeReadisEarnings
}

func ReadisEarningsNew() *TypeReadisEarnings {
	return &TypeReadisEarnings{}
}

func New() *TypeScheduler {
	return &TypeScheduler{
		config.GetConfig(),
		ReadisEarningsNew(),
	}
}
