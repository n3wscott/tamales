package main

import (
	"log"

	"github.com/n3wscott/tomles/pkg/commands"
)

func main() {
	cmds := commands.New()
	if err := cmds.Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
