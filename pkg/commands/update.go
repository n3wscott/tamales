package commands

import (
	"errors"
	"log"

	"github.com/spf13/cobra"

	"github.com/n3wscott/tomles/pkg/commands/options"
	"github.com/n3wscott/tomles/pkg/update"
)

func addUpdate(topLevel *cobra.Command) {
	vo := &options.VerboseOptions{}
	do := &options.DryRunOptions{}
	fo := &options.FilenameOptions{}
	co := &options.ConstraintOptions{}

	cmd := &cobra.Command{
		Use:       "update",
		ValidArgs: []string{},
		Short:     "Update changes the branch, version, or revision used for an import.",
		Example: `
  To change the dep on github.com/n3wscott/sockeye to the master branch in the local Gopkg.toml file,
  $ tomles update github.com/n3wscott/sockeye -b master -f ./Gopkg.toml

  Or the revision,
  $ tomles update github.com/n3wscott/sockeye -r 17fc779daec193476ff79fbe535ed00d426b08cb -f ./Gopkg.toml

  Or the version,
  $ tomles update github.com/n3wscott/sockeye -v v0.4.0 -f ./Gopkg.toml

  To see what the tomles command will produce,
  $ tomles update github.com/n3wscott/sockeye -b master -f ./Gopkg.toml --dry-run

  You can chain tomles output with a diff to see what will change,
  $ tomles update github.com/n3wscott/sockeye -b master -f ./Gopkg.toml --dry-run | diff ./Gopkg.toml -
`,

		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a import path argument")
			}
			co.Name = args[0]

			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			c := 0
			if co.Branch != "" {
				c++
			}
			if co.Version != "" {
				c++
			}
			if co.Revision != "" {
				c++
			}

			if c != 1 {
				return errors.New("command requires exactly one of: [branch, version, revision]")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Build up command.
			i := &update.Update{
				Filename: fo.Filename,
				DryRun:   do.DryRun,
				Verbose:  vo.Verbose,

				Name:     co.Name,
				Branch:   co.Branch,
				Version:  co.Version,
				Revision: co.Revision,
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
	options.AddConstraintsArgs(cmd, co)

	topLevel.AddCommand(cmd)
}
