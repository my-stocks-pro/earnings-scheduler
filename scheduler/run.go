package scheduler

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func (s *TypeScheduler) ApprovedTask() {

	errGET := s.HandleGET()
	if errGET != nil {
		fmt.Println(errGET)
	}

	errPOST := s.HandlePOST()
	if errPOST != nil {
		fmt.Println(errPOST)
	}

	fmt.Println(s.Requester.Response)
}

func (s *TypeScheduler) Run() {

	//TODO NEED use right time of schedule -> must change in conf

	fmt.Println("NewScheduler Create")

	approvedScheduler := gocron.NewScheduler()
	approvedScheduler.Every(s.Config.Tick).Seconds().Do(s.ApprovedTask)
	<-approvedScheduler.Start()
}
