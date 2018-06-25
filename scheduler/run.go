package scheduler

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func (s *TypeSchaduler) ApprovedTask() {

	data, errGET := s.HandleGET()
	if errGET != nil {
		fmt.Println(errGET)
	}

	resp, errPOST := s.HandlePOST(data)
	if errPOST != nil {
		fmt.Println(errPOST)
	}

	fmt.Println(resp)
}

func (s *TypeSchaduler) Run() {

	fmt.Println("NewScheduler Create")

	approvedScheduler := gocron.NewScheduler()
	approvedScheduler.Every(s.Config.Tick).Seconds().Do(s.ApprovedTask)
	<-approvedScheduler.Start()
}
