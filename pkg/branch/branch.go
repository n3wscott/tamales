package branch

import (
	"log"
)

type Branch struct {
	// TODO
	Filename string
	DryRun   bool
	Verbose  bool
}

func (b *Branch) Do() error {

	if b.Verbose {
		log.Println("Doing branch...", b.Filename)
	}

	return nil
}
