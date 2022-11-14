package main

import (
	"Cinnox-Homework/cmd"
	"fmt"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(fmt.Sprintf("cli Execute %v", err))
	}
}
