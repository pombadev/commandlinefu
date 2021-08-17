package main

import (
	"fmt"
	"os"
)

type App struct {
	urls *Urls
	cli  *Cli
}

func (app *App) Run() {
	if app.cli.previewThemes {
		app.PreviewThemes()
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
	} else {
		if len(app.cli.query) == 0 {
			fmt.Println("Please provide -query if -repl=false is set")
			os.Exit(1)
		}
	}
}

func NewApp() *App {
	urls := NewUrls()
	cli := NewCli()

	app := App{&urls, &cli}

	return &app
}
