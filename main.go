package main

import (
	"os"
)

const (
	AppName    string = "commandlinefu"
	AppVersion string = "v1.2.0"
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
