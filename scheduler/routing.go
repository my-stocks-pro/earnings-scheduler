package scheduler

import "github.com/gin-gonic/gin"

func (s *TypeScheduler) Routing() {

	s.Router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	s.Router.GET("/", func(c *gin.Context) {
		go s.Run()
	})

	s.Router.GET("/stop", func(c *gin.Context) {
		s.QuitRPC <- true
	})

}
