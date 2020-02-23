package commands

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "tomles",
		Short: "A tool to aid in yaml editing.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	addUpdate(cmds)

	return cmds
}
