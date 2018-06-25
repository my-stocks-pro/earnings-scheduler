package scheduler

import "net/http"

type TypeReadisEarnings struct {
	ID       string
	Count    string
	Earnings string
}

type TypeConfig struct {
	Tick    uint64
	URLGET  string
	URLPOST string
	RadisDB string
}

type TypeRequester struct {
	URL      string
	Method   string
	RadisDB  string
	Response *http.Response
	Body     []byte
}

type TypeScheduler struct {
	Config    *TypeConfig
	RedisData *TypeReadisEarnings
	Requester *TypeRequester
}

func NewRequester() *TypeRequester {
	return &TypeRequester{}
}

func ReadisEarningsNew() *TypeReadisEarnings {
	return &TypeReadisEarnings{}
}

func New() *TypeScheduler {
	return &TypeScheduler{
		GetConfig(),
		ReadisEarningsNew(),
		NewRequester(),
	}
}
