package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/n3wscott/tomles/pkg/commands"
)

func main() {
	cmds := &cobra.Command{
		Use:   "tomles",
		Short: "A tool to aid in yaml editing.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	commands.AddTomlesCommands(cmds)

	if err := cmds.Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
