package api

import (
	"Cinnox-Homework/cmd"
	"Cinnox-Homework/notify"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	conf   *cmd.Http
	http   *gin.Engine
	notify notify.INotify
}

func New(conf *cmd.Http, notify notify.INotify) *Server {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		conf:   conf,
		http:   gin.New(),
		notify: notify,
	}
}

func (s *Server) Run() error {
	s.Routes()
	return s.http.Run(s.conf.Addr)
}

func (s *Server) Routes() {
	s.http.POST("webhook", func(c *gin.Context) {
		if err := s.notify.Webhook(c.Request); err != nil {
			panic(err)
		}

		c.Status(http.StatusOK)
	})
}