package main

import (
	"Cinnox-Homework/cmd"
	"Cinnox-Homework/notify"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(fmt.Sprintf("cli Execute error: %v", err))
	}

	line, err := notify.New(cmd.Conf.Line)
	if err != nil {
		panic(line)
	}

	g := gin.New()

	g.POST("", func(c *gin.Context) {
		if err := line.ConsumeMessage(c.Request); err != nil {
			panic(err)
		}

		c.Status(http.StatusOK)
	})

	g.Run(cmd.Conf.Http.Addr)
}
