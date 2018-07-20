package scheduler

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func (s *TypeScheduler) EarningsTask() {

	data, errGET := s.Get()
	if errGET != nil {
		fmt.Println(errGET)
	}

	resp, errPOST := s.Post(data)
	if errPOST != nil {
		fmt.Println(errPOST)
	}

	fmt.Println(resp)
}

func (s *TypeScheduler) Run() {

	//TODO NEED use right time of schedule -> must change in conf

	fmt.Println("NewScheduler EarningsTask Create")

	approvedScheduler := gocron.NewScheduler()
	approvedScheduler.Every(s.Config.Tick).Seconds().Do(s.EarningsTask)
	<-approvedScheduler.Start()
}
