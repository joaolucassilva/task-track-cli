package main

import (
	"fmt"
	"github.com/joaolucassilva/task-track-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("err when execute this command %s", err)
		return
	}
}
