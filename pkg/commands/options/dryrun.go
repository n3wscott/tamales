package options

import "github.com/spf13/cobra"

// DryRunOptions
type DryRunOptions struct {
	DryRun bool
}

func AddDryRunArg(cmd *cobra.Command, po *DryRunOptions) {
	cmd.Flags().BoolVarP(&po.DryRun, "dry-run", "D", false,
		"Output what the new file would be.")
}
