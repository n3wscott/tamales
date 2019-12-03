package options

import (
	"github.com/spf13/cobra"
)

// FilenameOptions
type FilenameOptions struct {
	Filename string
}

func AddFilenameArg(cmd *cobra.Command, fo *FilenameOptions, required bool) {
	cmd.Flags().StringVarP(&fo.Filename, "filename", "f", "",
		"The path to the .toml file to use.")
	if required {
		if err := cmd.MarkFlagRequired("filename"); err != nil {
			panic(err)
		}
	}
}
