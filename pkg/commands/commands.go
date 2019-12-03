package commands

import (
	"github.com/spf13/cobra"
)

func AddTomlesCommands(topLevel *cobra.Command) {
	addBranch(topLevel)
}
