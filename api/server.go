package api

import (
	"Cinnox-Homework/cmd"
	"Cinnox-Homework/model"
	"Cinnox-Homework/notify"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	conf   *cmd.Http
	http   *gin.Engine
	notify notify.INotify
	model  *model.DB
}

func New(conf *cmd.Http, notify notify.INotify, model *model.DB) *Server {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		conf:   conf,
		http:   gin.New(),
		notify: notify,
		model:  model,
	}
}

func (s *Server) Run() error {
	s.Routes()
	return s.http.Run(s.conf.Addr)
}

func (s *Server) Routes() {
	s.http.POST("webhook", s.webhook)
	s.http.POST("send", s.sendMessage)
	s.http.POST("sends", s.sendsMessage)
	s.http.POST("broadcast", s.broadcastMessage)
}

func (s *Server) webhook(c *gin.Context) {
	if err := s.notify.Webhook(c.Request); err != nil {
		log.Printf("webhook api error: %v", err)
	}

	c.Status(http.StatusOK)
}

type sendMessageReq struct {
	UserId  string `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (s *Server) sendMessage(c *gin.Context) {
	var msg sendMessageReq
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "message or user_id field is required",
		})
		return
	}

	if err := s.notify.Send(msg.UserId, msg.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "send message bot api fail",
		})
		log.Printf("send message api error: %v", err)
		return
	}

	c.JSON(http.StatusOK, msg)
}

type sendsMessageReq struct {
	UserId  []string `json:"user_id" binding:"required"`
	Message string   `json:"message" binding:"required"`
}

func (s *Server) sendsMessage(c *gin.Context) {
	var msg sendsMessageReq
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "message or user_id field is required",
		})
		return
	}

	if err := s.notify.Sends(msg.UserId, msg.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "sends message bot api fail",
		})
		log.Printf("sends message api error: %v", err)
		return
	}

	c.JSON(http.StatusOK, msg)
}

type broadcastMessageReq struct {
	Message string `json:"message" binding:"required"`
}

func (s *Server) broadcastMessage(c *gin.Context) {
	var msg broadcastMessageReq
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "message field is required",
		})
		return
	}

	if err := s.notify.Broadcast(msg.Message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "broadcast message bot api fail",
		})
		log.Printf("broadcast message api error: %v", err)
		return
	}

	c.JSON(http.StatusOK, msg)
}
