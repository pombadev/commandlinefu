package main

import "os"

const (
	AppName    string = "commandlinefu"
	AppVersion string = "v2.0.0"
)

func main() {
	app := NewApp()

	if app.cli.listThemes {
		app.ListThemes()
		os.Exit(0)
	}

	if app.cli.version {
		app.Version()
		os.Exit(0)
	}

	if len(app.cli.query) >= 1 {
		app.Search()
		os.Exit(0)
	}

	if app.cli.repl {
		app.Repl()
	}
}
