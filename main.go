package main

import (
	"Cinnox-Homework/api"
	"Cinnox-Homework/cmd"
	"Cinnox-Homework/model"
	"Cinnox-Homework/notify"
	"context"
	"fmt"
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
}
