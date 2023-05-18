package main

import (
	"flag"
	"time"

	"github.com/nfielder/ts-infi-authkey/internal/cmd"
)

func main() {
	reusable := flag.Bool("reusable", false, "allocate a reusable authkey")
	ephemeral := flag.Bool("ephemeral", true, "allocate an ephemeral authkey")
	preauth := flag.Bool("preauth", true, "set the authkey as pre-authorised")
	tags := flag.String("tags", "", "comma-separated list of tags to apply to the authkey")
	expiry := flag.Duration("expiry", 5*time.Minute, "time until expiry of the authkey. Accepts string similar to 5m or 1h")
	flag.Parse()

	cmd.Run(cmd.CmdOpts{
		Reusable:  *reusable,
		Ephemeral: *ephemeral,
		Preauth:   *preauth,
		Tags:      *tags,
		Expiry:    *expiry,
	})
}
