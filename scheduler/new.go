package scheduler

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

type TypeScheduler struct {
	Config    *TypeConfig
	RedisData *TypeReadisEarnings
}

func ReadisEarningsNew() *TypeReadisEarnings {
	return &TypeReadisEarnings{}
}

func New() *TypeScheduler {
	return &TypeScheduler{
		GetConfig(),
		ReadisEarningsNew(),
	}
}
