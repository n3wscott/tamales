package commands

import (
	"errors"
	"log"

	"github.com/spf13/cobra"

	"github.com/n3wscott/tomles/pkg/branch"
	"github.com/n3wscott/tomles/pkg/commands/options"
)

func addBranch(topLevel *cobra.Command) {
	vo := &options.VerboseOptions{}
	do := &options.DryRunOptions{}
	fo := &options.FilenameOptions{}
	cmd := &cobra.Command{
		Use:   "branch",
		Short: "Branch changes the branch used for an import.",
		Example: `
  tomles branch something todo
`,
		Args: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Build up command.
			i := &branch.Branch{
				Filename: fo.Filename,
				DryRun:   do.DryRun,
				Verbose:  vo.Verbose,
			}

			// Run it.
			if err := i.Do(); err != nil {
				log.Fatalf("failed to run branch command: %v", err)
			}
		},
	}
	options.AddDryRunArg(cmd, do)
	options.AddVerboseArg(cmd, vo)
	options.AddFilenameArg(cmd, fo, true)

	topLevel.AddCommand(cmd)
}
