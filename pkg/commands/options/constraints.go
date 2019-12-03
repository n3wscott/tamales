package options

import (
	"github.com/spf13/cobra"
)

// ConstraintOptions
type ConstraintOptions struct {
	Name     string
	Branch   string
	Version  string
	Revision string
}

func AddConstraintsArgs(cmd *cobra.Command, co *ConstraintOptions) {
	cmd.Flags().StringVarP(&co.Branch, "branch", "b", "",
		"The value of the branch to use.")
	cmd.Flags().StringVarP(&co.Version, "version", "v", "",
		"The value of the version to use.")
	cmd.Flags().StringVarP(&co.Revision, "revision", "r", "",
		"The value of the revision to use.")
}
