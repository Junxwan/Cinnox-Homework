package main

import (
	"Cinnox-Homework/api"
	"Cinnox-Homework/cmd"
	"Cinnox-Homework/notify"
	"fmt"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(fmt.Sprintf("cli Execute error: %v", err))
	}

	line, err := notify.New(cmd.Conf.Line)
	if err != nil {
		panic(err)
	}

	server := api.New(&cmd.Conf.Http, line)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
