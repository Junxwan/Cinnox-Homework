package main

import (
	"Cinnox-Homework/api"
	"Cinnox-Homework/cmd"
	"Cinnox-Homework/model"
	"Cinnox-Homework/notify"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(fmt.Sprintf("cli Execute error: %v", err))
	}

	ctx := context.Background()

	db, err := model.NewDB(ctx, cmd.Conf.Databases)
	if err != nil {
		panic(err)
	}

	line, err := notify.New(cmd.Conf.Line, db)
	if err != nil {
		panic(err)
	}

	server := api.New(&cmd.Conf.Http, line, db)
	if err := server.Run(); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Printf("goim-comet get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if err := db.Close(context.Background()); err != nil {
				log.Printf("close mongodb error: %v", err)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
