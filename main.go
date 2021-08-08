package main

import (
	"os"
)

var (
	AppRevision string
	AppVersion  string
	AppName     string = "commandlinefu"
)

func main() {
	cli := NewCli()

	if cli.version {
		cli.Version()
		os.Exit(1)
	}

	if len(cli.query) >= 1 {
		cli.Search()
		os.Exit(0)
	}

	if cli.repl {
		cli.Repl()
	}
}
