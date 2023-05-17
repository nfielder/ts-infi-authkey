package main

import (
	"flag"

	"github.com/nfielder/ts-infi-authkey/internal/cmd"
)

func main() {
	reusable := flag.Bool("reusable", false, "allocate a reusable authkey")
	ephemeral := flag.Bool("ephemeral", true, "allocate an ephemeral authkey")
	preauth := flag.Bool("preauth", true, "set the authkey as pre-authorised")
	tags := flag.String("tags", "", "comma-separated list of tags to apply to the authkey")
	flag.Parse()

	cmd.Run(cmd.CmdOpts{
		Reusable:  *reusable,
		Ephemeral: *ephemeral,
		Preauth:   *preauth,
		Tags:      *tags,
	})
}
