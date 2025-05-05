package main

import (
	"context"

	_ "github.com/warpstreamlabs/bento/public/components/all"
	"github.com/warpstreamlabs/bento/public/service"

	// import your plugins:
	_ "plugin_testbench/processors"
)

func main() {
	// RunCLI accepts a number of optional functions:
	// https://pkg.go.dev/github.com/warpstreamlabs/bento/public/service#CLIOptFunc
	service.RunCLI(context.Background())
}
